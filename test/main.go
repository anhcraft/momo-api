package main

import (
	"fmt"
	"momo-api/api/auth"
	"momo-api/api/model"
	"momo-api/api/trans"
	"momo-api/utils"
	"os"
)

func login() {
	phone := utils.RequireEnvSetting("test_phone")
	pass := utils.RequireEnvSetting("test_pass")
	u := model.CreateStandardUser(phone, "Guest")
	fmt.Println("Verifying phone...")
	err, _ := auth.VerifyPhone(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Requesting OTP...")
	err = auth.RequestOTP(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Please enter OTP code: ")
	var otp string
	_, err = fmt.Scanln(&otp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Verifying OTP...")
	err = auth.VerifyOTP(u, otp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Logging in...")
	err = auth.Login(u, pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Create("./test/user.json")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(u.ToJson())
}

func getTrans() {
	data, err := os.ReadFile("./test/user.json")
	if err != nil {
		return
	}
	u, _ := model.ParseUser(string(data))
	it := trans.BrowseDefaultRecentTransactions(u)
	it.Next()
	for _, v := range it.Buffer {
		fmt.Printf("ID %s: %s sent %.2f to %s\n", v.Id, v.Sender.Name, v.Amount, v.Receiver.Name)
	}
}

func getTransId() {
	data, err := os.ReadFile("./test/user.json")
	if err != nil {
		return
	}
	u, _ := model.ParseUser(string(data))
	fmt.Print("Please enter trans id: ")
	var tid string
	_, err = fmt.Scanln(&tid)
	if err != nil {
		fmt.Println(err)
		return
	}
	v := &model.P2PTransaction{Id: tid}
	err = trans.GetTransactionDetail(u, v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ID %s: %s sent %.2f to %s (msg: %s)\n", v.Id, v.Sender.Name, v.Amount, v.Receiver.Name, v.Message)
}

func sendMoney() {
	data, err := os.ReadFile("./test/user.json")
	if err != nil {
		return
	}
	u, _ := model.ParseUser(string(data))
	t := &model.P2PTransaction{
		Id:       "",
		Amount:   2000,
		Message:  "Gui tien",
		Sender:   u,
		Receiver: model.CreatePartner(utils.RequireEnvSetting("receiver_phone"), utils.RequireEnvSetting("receiver_name")),
		Time:     0,
	}
	r, err := trans.InitTransfer(t)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = trans.ConfirmTransfer(t, r)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func showMenu() {
	fmt.Println()
	fmt.Println("=== MENU ===")
	fmt.Println("- 0 to exit")
	fmt.Println("- 1 to login")
	fmt.Println("- 2 to get trans")
	fmt.Println("- 3 to get trans id")
	fmt.Println("- 4 to send money")
	fmt.Print("Type your option: ")
}

func main() {
	for true {
		showMenu()
		var n int
		_, err := fmt.Scanln(&n)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		switch n {
		case 0:
			{
				os.Exit(0)
			}
		case 1:
			{
				login()
				break
			}
		case 2:
			{
				getTrans()
				break
			}
		case 3:
			{
				getTransId()
				break
			}
		case 4:
			{
				sendMoney()
				break
			}
		}
	}
}
