package message

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"momo-api/api/model"
	"momo-api/utils"
	"momo-api/utils/crypto"
	"net/http"
	"strconv"
	"strings"
)

type RequestBuilder struct {
	Headers     map[string]string
	Body        string
	URL         string
	EncryptBody bool
	Logging     bool
}

func CreateRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		Headers: make(map[string]string),
		Body:    "",
		Logging: cast.ToBool(utils.GetEnvSetting("request_logging", "false")),
	}
}

func (r *RequestBuilder) SetURL(endpoint string, subPath string) {
	r.URL = endpoint + subPath
}

func (r *RequestBuilder) SetBodyJSON(object any) error {
	b, err := json.Marshal(object)
	if err == nil {
		r.Body = string(b)
	}
	return err
}

func (r *RequestBuilder) SetHeader(key string, val string) {
	r.Headers[key] = val
}

func (r *RequestBuilder) EnableBodyEncryption() {
	r.EncryptBody = true
}

// vn.momo.core.modules.networking.http.utilities.cipher.RestRequestEncryptor
// vn.momo.core.modules.networking.http.request.decorators.EncryptRequestBodyDecorator
// vn.momo.core.modules.networking.http.response.DecryptResponseConverter
func (r *RequestBuilder) Fetch(requester *model.User) ([]byte, error) {
	requestKeyRaw := []byte(utils.RandString(32, utils.Alphanumeric))
	content := []byte(r.Body)
	if r.EncryptBody {
		n, err := crypto.EncryptAes256CbcPKCS7ToBase64(content, requestKeyRaw)
		if err != nil {
			return nil, err
		}
		content = []byte(n)
	}

	req, err := http.NewRequest(http.MethodPost, r.URL, bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	r.SetHeader("Content-Type", "application/json")
	r.SetHeader("Accept", "application/json")
	r.SetHeader("Accept-Charset", "UTF-8")
	r.SetHeader("user-Agent", utils.UserAgent)
	r.SetHeader("device_os", strings.ToUpper(utils.DeviceOS))
	r.SetHeader("app_version", strconv.Itoa(utils.AppVersion))
	r.SetHeader("app_code", utils.AppCode)
	r.SetHeader("channel", utils.Channel)
	r.SetHeader("lang", utils.LangCode)

	if requester != nil {
		r.SetHeader("user_phone", requester.Phone)
		r.SetHeader("userId", requester.Phone)
		if requester.Profile != nil {
			if r.EncryptBody {
				requestKey, err1 := crypto.EncryptRSAToBase64(requestKeyRaw, []byte(requester.Profile.RequestEncryptKey))
				if err1 != nil {
					return nil, err1
				}
				r.SetHeader("requestKey", requestKey)
			}

			r.SetHeader("momo-session-key-tracking", requester.Profile.SessionKeyTracking)
			if requester.Profile.AuthToken != "" {
				r.SetHeader("authorization", "Bearer "+requester.Profile.AuthToken)
			}
			r.SetHeader("sessionKey", requester.Profile.SessionKey)
			r.SetHeader("agent_id", requester.Profile.AgentId)
		}
	}
	for k, v := range r.Headers {
		req.Header.Add(k, strings.TrimSpace(v))
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err = res.Body.Close(); err != nil {
		return nil, err
	}

	if r.Logging {
		fmt.Println(strings.Repeat("=", 30))
		fmt.Printf("%s [Encrypted: %v]\n", r.URL, r.EncryptBody)
		for k, _ := range req.Header {
			fmt.Printf("- %s: %s\n", k, req.Header.Get(k))
		}
		fmt.Println(r.Body)
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println(string(body))
		fmt.Println(strings.Repeat("=", 30))
	}

	if r.EncryptBody && len(body) > 0 {
		body, err = crypto.DecryptAes256CbcPKCS7FromBase64(string(body), requestKeyRaw)
		if err != nil {
			return nil, err
		}
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	return body, nil
}
