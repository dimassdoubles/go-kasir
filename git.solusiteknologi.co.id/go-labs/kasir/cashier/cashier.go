package cashier

import (
	"errors"
	"fmt"
	"strings"

	"git.solusiteknologi.co.id/go-labs/kasir/cstmutil"
	"git.solusiteknologi.co.id/go-labs/kasir/dao/userdao"
	"git.solusiteknologi.co.id/go-labs/kasir/errormsg"
	"git.solusiteknologi.co.id/go-labs/kasir/tables"
	"github.com/jackc/pgx/v4"
)

type Cashier struct {
	user tables.LearnUser
	Tx   pgx.Tx
}

func (c *Cashier) Start() error {
	err := c.login()
	if err != nil {
		cstmutil.GiveSomeSpace()
		return err
	}

	cstmutil.GiveSomeSpace()
	option, err := c.selectMenu()
	if err != nil {
		cstmutil.GiveSomeSpace()
		return err
	}

	cstmutil.GiveSomeSpace()
	err = c.processMenu(option)
	if err != nil {
		cstmutil.GiveSomeSpace()
		return err
	}
	return nil
}

func (c *Cashier) login() error {
	fmt.Println("=====================================================")
	fmt.Println("*** Silakan Login Untuk Menggunakan Program Kasir ***")
	fmt.Println("-----------------------------------------------------")

	var username string

	fmt.Print("User     : ")
	_, err := fmt.Scanln(&username)
	if err != nil {
		fmt.Println(err)
	}

	var password string
	fmt.Print("Password : ")
	_, err = fmt.Scanln(&password)
	if err != nil {
		fmt.Println(err)
	}

	// menghapus spasi di awal dan akhir string
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	user, err := userdao.Login(userdao.InputLogin{
		Tx:       c.Tx,
		Username: username,
		Password: password,
	})

	if err != nil {
		return errors.New(errormsg.GAGAL_LOGIN)
	}

	c.user = *user
	return nil
}

func (c Cashier) selectMenu() (string, error) {
	var option string

	fmt.Println("Silakan pilih menu")
	fmt.Println("1. Penjualan Barang")
	fmt.Println("2. Lihat Barang")
	fmt.Println("3. Tambah Barang")
	fmt.Println("4. Riwayat Penjualan")
	fmt.Println("5. Logout")
	fmt.Println()
	fmt.Print("Input pilihan 1 - 3 : ")

	_, err := fmt.Scanln(&option)
	if err != nil {
		cstmutil.GiveSomeSpace()
		fmt.Println(err)
		cstmutil.GiveSomeSpace()
		return c.selectMenu()
	} else {
		option = strings.TrimSpace(option)
		if isValidOption(option) {
			return option, nil
		}
		cstmutil.GiveSomeSpace()
		fmt.Println(errormsg.PILIHAN_TIDAK_DAPAT_DIPROSES)
		cstmutil.GiveSomeSpace()
		return c.selectMenu()
	}
}

func isValidOption(option string) bool {
	options := []string{"1", "2", "3", "4", "5"}

	for _, v := range options {
		if v == option {
			return true
		}
	}

	return false
}

func (c Cashier) processMenu(option string) error {
	switch option {
	case "1":
		fmt.Println("penjualan barang")
	case "2":
		fmt.Println("lihat barang")
	case "3":
		fmt.Println("tambah barang")
	case "4":
		fmt.Println("riwayat penjualan")
	case "5":
		fmt.Println("logout")
	default:
		fmt.Println(errormsg.PILIHAN_TIDAK_DAPAT_DIPROSES)
	}

	return nil
}
