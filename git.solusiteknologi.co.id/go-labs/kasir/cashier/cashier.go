package cashier

import (
	"errors"
	"fmt"
	"strings"

	"git.solusiteknologi.co.id/go-labs/kasir/cstmutil"
	"git.solusiteknologi.co.id/go-labs/kasir/dao/penjualandao"
	"git.solusiteknologi.co.id/go-labs/kasir/dao/penjualanitemdao"
	"git.solusiteknologi.co.id/go-labs/kasir/dao/productdao"
	"git.solusiteknologi.co.id/go-labs/kasir/dao/userdao"
	"git.solusiteknologi.co.id/go-labs/kasir/errormsg"
	"git.solusiteknologi.co.id/go-labs/kasir/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glutil"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
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
		c.doPenjualan()
	case "2":
		c.doLihatBarang()
	case "3":
		fmt.Println("riwayat penjualan")
	case "4":
		fmt.Println("riwayat penjualan")
	default:
		fmt.Println(errormsg.PILIHAN_TIDAK_DAPAT_DIPROSES)
	}

	return nil
}

func (c Cashier) doPenjualan() error {
	fmt.Println("=============================")
	fmt.Println("** Menu Penjualan Barang **")
	fmt.Println("—----------------------------")
	cstmutil.GiveSomeSpace()

	fmt.Println("Pilih Produk")

	productList, err := productdao.GetProducts(productdao.InputGetProducts{
		Tx: c.Tx,
	})

	if err != nil {
		return err
	}

	penjualanItems := []PenjualanItem{}

	penjualanItems = append(penjualanItems, inputPenjualanItem(productList))
	showPenjualanItem(penjualanItems)
	cstmutil.GiveSomeSpace()

	for isSelectNextProduct() {
		cstmutil.GiveSomeSpace()
		penjualanItems = append(penjualanItems, inputPenjualanItem(productList))
		cstmutil.GiveSomeSpace()
		showPenjualanItem(penjualanItems)
		cstmutil.GiveSomeSpace()
	}

	cstmutil.GiveSomeSpace()
	totalPenjualan := getTotalPenjualan(penjualanItems)
	pembayaran := inputPembayaran(penjualanItems)
	kembalian := pembayaran.Sub(totalPenjualan)

	auditUserId := c.user.UserId
	auditDatetime := glutil.DateTimeNow()

	penjualan, err := penjualandao.Add(penjualandao.InputAdd{
		Tx:              c.Tx,
		TotalPenjualan:  totalPenjualan,
		TotalPembayaran: pembayaran,
		TotalKembalian:  kembalian,
		AuditUserId:     auditUserId,
		AuditDatetime:   auditDatetime,
	})

	if err != nil {
		return err
	}

	for _, penjualanItem := range penjualanItems {
		_, err := penjualanitemdao.Add(penjualanitemdao.InputAdd{
			Tx:            c.Tx,
			PenjualanId:   penjualan.PenjualanId,
			ProductId:     penjualanItem.Product.ProductId,
			Qty:           penjualanItem.Qty,
			Price:         penjualanItem.Product.Price,
			AuditUserId:   auditUserId,
			AuditDatetime: auditDatetime,
		})

		if err != nil {
			return err
		}
	}

	fmt.Println("=======================")
	fmt.Println("** Summary belanja **")
	fmt.Println("—----------------------------------------")
	fmt.Println("1. Total belanja         ", totalPenjualan)
	fmt.Println("2. Pembayaran            ", pembayaran)
	fmt.Println("3. Kembalian             ", kembalian)

	cstmutil.GiveSomeSpace()
	fmt.Println("Terimakasih Telah Berbelanja ^o^")

	return nil
}

type PenjualanItem struct {
	Product tables.LearnProduct
	Qty     decimal.Decimal
}

func inputPembayaran(penjualanItems []PenjualanItem) decimal.Decimal {
	var pembayaran int64

	fmt.Print("Masukan Pembayaran : ")
	_, err := fmt.Scan(&pembayaran)

	if err != nil {
		cstmutil.GiveSomeSpace()
		fmt.Println(err)

		cstmutil.GiveSomeSpace()
		return inputPembayaran(penjualanItems)
	}

	if pembayaran < getTotalPenjualan(penjualanItems).IntPart() {
		cstmutil.GiveSomeSpace()
		fmt.Println(errormsg.PEMBAYARAN_KURANG)

		cstmutil.GiveSomeSpace()
		return inputPembayaran(penjualanItems)
	}

	return decimal.NewFromInt(pembayaran)
}

func isSelectNextProduct() bool {
	var option string
	fmt.Print("Tambah Produk Lagi Y / N ? ")
	_, err := fmt.Scanln(&option)

	if err == nil {
		if strings.TrimSpace(option) == "Y" {
			return true
		}
	}

	return false
}

func showPenjualanItem(penjualanItems []PenjualanItem) {
	cstmutil.GiveSomeSpace()

	fmt.Println("Total Pembelian   ", getTotalPenjualan(penjualanItems))
	fmt.Println("—-------------------------")
	for i, v := range penjualanItems {
		fmt.Println(i+1, ".", v.Product.ProductCode, v.Product.ProductName, v.Product.Price.Mul(v.Qty))
	}

}

func getTotalPenjualan(penjualanItems []PenjualanItem) decimal.Decimal {
	total := decimal.NewFromInt(0)
	for _, v := range penjualanItems {
		total = total.Add(v.Product.Price.Mul(v.Qty))
	}

	return total
}

func inputPenjualanItem(productList []*tables.LearnProduct) PenjualanItem {
	showProduct(productList)

	product := selectProduct(productList)
	qty := inputQty()

	return PenjualanItem{
		Product: product,
		Qty:     qty,
	}
}

func showProduct(productList []*tables.LearnProduct) {
	i := 1
	for _, v := range productList {
		fmt.Println(i, ". ", v)
		i++
	}
}

func selectProduct(productList []*tables.LearnProduct) tables.LearnProduct {
	cstmutil.GiveSomeSpace()
	var option int
	fmt.Print("Pilih produk 1 - ", len(productList), " : ")
	_, err := fmt.Scan(&option)

	if err != nil {
		cstmutil.GiveSomeSpace()
		fmt.Println(err)

		cstmutil.GiveSomeSpace()
		showProduct(productList)
		return selectProduct(productList)
	}

	if option <= 0 || option > len(productList) {
		cstmutil.GiveSomeSpace()
		fmt.Println(errormsg.PILIHAN_TIDAK_DAPAT_DIPROSES)

		cstmutil.GiveSomeSpace()
		showProduct(productList)
		return selectProduct(productList)
	}

	return *productList[option-1]
}

func inputQty() decimal.Decimal {
	var qty int64

	cstmutil.GiveSomeSpace()
	fmt.Print("Ketik jumlah pembelian : ")
	_, err := fmt.Scanln(&qty)

	if err != nil {
		cstmutil.GiveSomeSpace()
		fmt.Println(err)

		return inputQty()
	}

	if qty <= 0 {
		cstmutil.GiveSomeSpace()
		fmt.Println(errormsg.MINIMAL_JUMLAH_1)

		return inputQty()
	}

	return decimal.NewFromInt(qty)
}

func (c Cashier) doLihatBarang() error {
	cstmutil.GiveSomeSpace()
	fmt.Println("=============================")
	fmt.Println("** Menu Lihat Barang **")
	fmt.Println("—----------------------------")

	learnProducts, err := productdao.GetProducts(productdao.InputGetProducts{
		Tx: c.Tx,
	})

	if err != nil {
		return err
	}

	showProduct(learnProducts)

	c.doAddNewProduct()

	return nil

}

func isAddNewProduct() bool {
	var option string
	fmt.Print("Tambah Baru? Y / N : ")
	_, err := fmt.Scanln(&option)

	if err == nil {
		if strings.TrimSpace(option) == "Y" {
			return true
		}
	}

	return false
}

func (c Cashier) doAddNewProduct() error {
	products := []*tables.LearnProduct{}

	for isAddNewProduct() {
		product, err := productdao.Add(productdao.InputAdd{
			Tx:            c.Tx,
			ProductCode:   inputKodeBarang(),
			ProductName:   inputNamaBarang(),
			Price:         inputHargaBarang(),
			AuditDatetime: glutil.DateTimeNow(),
			AuditUserId:   c.user.UserId,
		})

		if err != nil {
			cstmutil.GiveSomeSpace()
			fmt.Println(errormsg.GAGAL_TAMBAH_PRODUK)
			continue
		}

		products = append(products, product)

		cstmutil.GiveSomeSpace()
		fmt.Println("Barang Sukses Ditambahkan")
		showProduct(products)

	}

	return nil
}

func inputKodeBarang() string {
	cstmutil.GiveSomeSpace()
	var kodeBarang string
	fmt.Print("Masukkan Kode Barang  : ")
	_, err := fmt.Scan(&kodeBarang)
	if err != nil {
		cstmutil.GiveSomeSpace()
		fmt.Println(err)
		return inputKodeBarang()
	}

	if len(kodeBarang) > 4 || len(kodeBarang) < 4 {
		cstmutil.GiveSomeSpace()
		fmt.Println(errormsg.KODE_BARANG_TIDAK_VALID)
		return inputKodeBarang()
	}

	return kodeBarang

}

func inputNamaBarang() string {
	var namaBarang string
	fmt.Print("Masukkan Nama Barang  : ")
	_, err := fmt.Scan(&namaBarang)
	if err != nil {
		cstmutil.GiveSomeSpace()
		fmt.Println(err)
		return inputNamaBarang()
	}

	return namaBarang

}

func inputHargaBarang() decimal.Decimal {
	var hargaBarang int64
	fmt.Print("Masukkan Harga Barang : ")
	_, err := fmt.Scan(&hargaBarang)
	if err != nil {
		cstmutil.GiveSomeSpace()
		fmt.Println(err)
		return inputHargaBarang()
	}

	if hargaBarang <= 0 {
		cstmutil.GiveSomeSpace()
		fmt.Println(errormsg.MINIMAL_JUMLAH_1)
		return inputHargaBarang()
	}

	return decimal.NewFromInt(hargaBarang)

}
