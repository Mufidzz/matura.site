package main

import (
	"../main/import/restful"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var ty_users restful.User
var users []restful.User
var response restful.Response

var stInstansi restful.Instansi
var arInstansi []restful.Instansi

var templates *template.Template

func main() {
	router := mux.NewRouter()
	templates = template.Must(template.ParseGlob("templates/*.html"))

	fs := http.FileServer(http.Dir("./templates/static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	//router.HandleFunc("/users/{id}", getUser).Methods("VIEW")

	router.HandleFunc("/users", getUsers).Methods("VIEW")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users", deleteUser).Methods("DELETE")
	router.HandleFunc("/users", updateUser).Methods("PUT")

	router.HandleFunc("/register/addr/province", getProvince).Methods("VIEW")
	router.HandleFunc("/register/addr/kabupaten/{id}", getKabupaten).Methods("VIEW")
	router.HandleFunc("/register/addr/kecamatan/{id}", getKecamatan).Methods("VIEW")

	router.HandleFunc("/addr/univ/{table}", getUniveristy).Methods("GET")

	router.HandleFunc("/", indexGetHandler).Methods("GET")
	router.HandleFunc("/register", registerGetHandler).Methods("GET")
	router.HandleFunc("/admin", adminGetHandler).Methods("GET")
	router.HandleFunc("/user", userGetHandler).Methods("GET")
	router.HandleFunc("/try", trialGetHandler).Methods("GET")

	http.Handle("/", router)
	fmt.Println("Connected to port 1174")
	log.Fatal(http.ListenAndServe(":1174", router))
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	templates.ExecuteTemplate(w, "index.html", nil)
}

func userGetHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	templates.ExecuteTemplate(w, "user2.html", nil)
}

func registerGetHandler(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	templates.ExecuteTemplate(w, "register.html", nil)
}

func trialGetHandler(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	templates.ExecuteTemplate(w, "try.html", nil)
}

func adminGetHandler(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	var data []string

	param := r.URL.Query()
	if _, ok := param["account"]; ok {
		data = []string{"Manage User","Account Logs","User Registered : 192746"}
	} else if _, ok := param["news"]; ok {
		data = []string{"Manage News","Create News","New Added : 1946"}
	} else if _, ok := param["shop"]; ok {
		data = []string{"Manage Product","Create Product","Report","Total Product : 765"}
	} else if _, ok := param["data"]; ok {
		data = []string{"Register Data","Shop Data"}
	}
	templates.ExecuteTemplate(w, "admin.html",data)

}


func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/matura_db.alumni")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func instansiConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/matura_db.instansi")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func daerahConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/matura_db.daerah")

	if err != nil {
		log.Fatal(err)
	}

	return db
}


func getUniveristy(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	db := instansiConnect()
	defer db.Close()
	urlParam := mux.Vars(r)

	rows, err := db.Query("SELECT Instansi FROM `tb."+ urlParam["table"] + "` ORDER BY Instansi ASC")

	if err != nil {
		log.Print(err)
	}

	for rows.Next(){
		if err := rows.Scan(&stInstansi.Instansi); err != nil{
			log.Fatal(err.Error())
		} else {
			arInstansi = append(arInstansi, stInstansi)
		}
	}



	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(arInstansi)
}

func getProvince(w http.ResponseWriter, r *http.Request) {
	var arProvinsi []restful.Provinsi
	var stProvinsi restful.Provinsi

	db := daerahConnect()
	defer db.Close()

	rows, err := db.Query("SELECT id_prov, nama FROM `provinsi` ORDER BY nama ASC")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&stProvinsi.Id, &stProvinsi.Nama); err != nil{
			log.Fatal(err)
		} else {
			arProvinsi = append(arProvinsi, stProvinsi)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin","*")
	json.NewEncoder(w).Encode(arProvinsi)
}

func getKabupaten(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var arKab []restful.Kabupaten
	var stKab restful.Kabupaten
	urlParam := mux.Vars(r)

	db := daerahConnect()
	defer db.Close()

	rows, err := db.Query("SELECT id_kab, nama FROM `kabupaten` WHERE id_prov = ? ORDER BY nama ASC", urlParam["id"])

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&stKab.Id, &stKab.Nama); err != nil{
			log.Fatal(err)
		} else {
			arKab = append(arKab, stKab)
		}
	}


	log.Print("Have Some good request")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(arKab)

}

func getKecamatan(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var arKec []restful.Kecamatan
	var stKec restful.Kecamatan
	urlParam := mux.Vars(r)

	db := daerahConnect()
	defer db.Close()

	rows, err := db.Query("SELECT id_kec, nama FROM `kecamatan` WHERE id_kab = ? ORDER BY nama ASC", urlParam["id"])

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&stKec.Id, &stKec.Nama); err != nil{
			log.Fatal(err)
		} else {
			arKec = append(arKec, stKec)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(arKec)

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	db := Connect()
	defer db.Close()

	rows, err := db.Query("SELECT AID,Nama,TempatLahir,TanggalLahir,AlamatAsal,AlamatTinggal,TahunLulus FROM `tb.biodata`")

	if err != nil {
		log.Print(err)
	}

	for rows.Next(){
		if err := rows.Scan(&ty_users.AID, &ty_users.Nama, &ty_users.POB, &ty_users.DOB, &ty_users.AlamatA, &ty_users.AlamatB, &ty_users.TahunLulus); err != nil{
			log.Fatal(err.Error())
		} else {
			users = append(users, ty_users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = users
	response.Ins = nil

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func createUser(w http.ResponseWriter, r *http.Request){
enableCors(&w)
	db:= Connect()
	defer db.Close()

	err := json.NewDecoder(r.Body).Decode(&ty_users)

	if err != nil {
		panic(err)
	}

	phash, err := bcrypt.GenerateFromPassword([]byte(ty_users.PwordHash),bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	id		:= ty_users.ID
	pwd		:= phash

	nama 	:= ty_users.Nama
	pob 	:= ty_users.POB
	dob 	:= ty_users.DOB
	alamata	:= ty_users.AlamatA
	alamatb	:= ty_users.AlamatB
	thlulus	:= ty_users.TahunLulus

	status	:= ty_users.Status
	instansi:= ty_users.Instansi

	facebook:= ty_users.Facebook
	line 	:= ty_users.Line
	wa 		:= ty_users.Whatsapp

	aid 	:= thlulus + generateAID()

	_, err = db.Exec("INSERT INTO `tb.account` (AID, Username, Password) VALUES (?, ?, ?)",
		aid,
		id,
		pwd,
		)

	_, err = db.Exec("INSERT INTO `tb.biodata` (`AID`, `Nama`, `TempatLahir`, `TanggalLahir`, `AlamatAsal`, `AlamatTinggal`, `TahunLulus`) VALUES (?, ?, ?, ?, ?, ?, ?)",
		aid,
		nama,
		pob,
		dob,
		alamata,
		alamatb,
		thlulus)

	_, err = db.Exec("INSERT INTO `tb.instansi` (`AID`, `Status`, `Instansi`) VALUES (?, ?, ?)",
		aid,
		status,
		instansi)

	_, err = db.Exec("INSERT INTO `tb.media` (`AID`, `Facebook`, `Line`, `Whatsapp`) VALUES (?, ?, ?, ?)",
		aid,
		facebook,
		line,
		wa)

	if err != nil {
		log.Print(err)
	} else {
		response.Status = 1
		response.Message = "Success Add"
		log.Print("New data to database : \n\tAID : " + aid + "\n\tName :" + nama)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func deleteUser(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	db:= Connect()
	defer db.Close()

	err := json.NewDecoder(r.Body).Decode(&ty_users)
	if err != nil {
		panic(err)
	}

	aid 	:= ty_users.AID

	_, err = db.Exec("DELETE FROM `tb.media` WHERE `tb.media`.`AID` = ?",
		aid,
	)

	_, err = db.Exec("DELETE FROM `tb.instansi` WHERE `tb.instansi`.`AID` = ?",
		aid,
	)

	_, err = db.Exec("DELETE FROM `tb.biodata` WHERE `tb.biodata`.`AID` = ?",
		aid,
	)

	_, err = db.Exec("DELETE FROM `tb.account` WHERE `tb.account`.`AID` = ?",
		aid,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Deleted"
	log.Print("Delete Data With AID :" +aid)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func updateUser(w http.ResponseWriter, r *http.Request)  {
	enableCors(&w)
	db:= Connect()
	defer db.Close()

	err := json.NewDecoder(r.Body).Decode(&ty_users)
	if err != nil {
		panic(err)
	}



	aid 	:= ty_users.AID
	id		:= ty_users.ID
	pwd		:= ty_users.PwordHash

	nama 	:= ty_users.Nama
	pob 	:= ty_users.POB
	dob 	:= ty_users.DOB
	alamata	:= ty_users.AlamatA
	alamatb	:= ty_users.AlamatB
	thlulus	:= ty_users.TahunLulus

	status	:= ty_users.Status
	instansi:= ty_users.Instansi

	facebook:= ty_users.Facebook
	line 	:= ty_users.Line
	wa 		:= ty_users.Whatsapp

	_, err = db.Exec("UPDATE `tb.account` SET `Username` = ?, `Password` = ? WHERE `tb.account`.`AID` = ?",
		id,
		pwd,
		aid)

	_, err = db.Exec("UPDATE `tb.biodata` SET `Nama` = ?, `TempatLahir` = ?, `TanggalLahir` = ?, `AlamatAsal` = ?, `AlamatTinggal` = ?, `TahunLulus` = ? WHERE `tb.biodata`.`AID` = ?",
		nama,
		pob,
		dob,
		alamata,
		alamatb,
		thlulus,
		aid)
	_, err = db.Exec("UPDATE `tb.instansi` SET `Status` = ?,`Instansi` = ? WHERE `tb.instansi`.`AID` = ?",
		status,
		instansi,
		aid)

	_, err = db.Exec("UPDATE `tb.media` SET `Facebook` = ?, `Line` = ?, `Whatsapp` = ? WHERE `tb.media`.`AID` = ?",

		facebook,
		line,
		wa,
		aid)


	response.Status = 1
	response.Message = "Update Success"
	log.Print("Succesfully Update AID : "+aid)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func generateAID() string{
	db := Connect()
	defer db.Close()

	thlulus	:= ty_users.TahunLulus

	rows, err := db.Query("SELECT AID FROM `tb.biodata` WHERE `TahunLulus` = ? ORDER BY AID DESC LIMIT 1",
		thlulus)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		e := rows.Scan(&ty_users.AID)

		if e != nil {
			log.Fatal(e)
		}
	}

	if len(ty_users.AID) <=4 {
		ty_users.AID += "10000"
	}

	lf := string(ty_users.AID[len(ty_users.AID)-5:])
	m, _ := strconv.Atoi(lf)

	return strconv.Itoa(m+1)
}
