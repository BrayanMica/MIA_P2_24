package Analyzer

import (
	"MIA_P1_201907343/DiskManagement"
	"MIA_P1_201907343/FileSystem"
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var re = regexp.MustCompile(`-(\w+)=("[^"]+"|\S+)`)

type responseList struct {
	Status int64    `json:"Status"`
	List   []string `json:"List"`
}

type responseString struct {
	Status int64  `json:"Status"`
	Value  string `json:"Value"`
}

type loginValues struct {
	User     string `json:"User"`
	Password string `json:"Password"`
}

var newResponseList responseList
var continuar bool = false

func addToResponseList() {
	newResponseList.List = append(newResponseList.List, "A.disk")
	newResponseList.List = append(newResponseList.List, "B.disk")
	newResponseList.List = append(newResponseList.List, "C.disk")
}

func postMethod(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Data")
	}
	var newLoginValues loginValues
	json.Unmarshal(reqBody, &newLoginValues)

	fmt.Println(newLoginValues.User)
	fmt.Println(newLoginValues.Password)

	newResponseList.Status = 200
	newResponseList.List = []string{}

	addToResponseList()

	for _, item := range newResponseList.List {
		fmt.Println(item)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newResponseList)
}

func getMethod(w http.ResponseWriter, r *http.Request) {
	var newResponseString responseString
	newResponseString.Status = 200
	newResponseString.Value = "Si funciona el get XD"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newResponseString)
}

func handleRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Weltome to my  API :D")
}

func getCommandAndParams(input string) (string, string) {
	parts := strings.Fields(input)
	if len(parts) > 0 {
		command := strings.ToLower(parts[0])
		params := strings.Join(parts[1:], " ")
		return command, params
	}
	return "", input
}

// Funcion para agregar discos a la lista post
func AgreagarDiscos() {
	newResponseList.Status = 200
	newResponseList.List = append(newResponseList.List, "Disco 1")
	newResponseList.List = append(newResponseList.List, "Disco 2")
	newResponseList.List = append(newResponseList.List, "Disco 3")
}

func Analyze() {
	for {
		var input string
		fmt.Print("-> ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()

		command, params := getCommandAndParams(input)

		//fmt.Println("comando: ", command, "Parametros: ", params)

		AnalyzeCommnad(command, params)

		//mkdisk -size=3000 -unit=K -fit=BF
		//fdisk -size=300 -driveletter=A -name=Particion1
		//mount -driveletter=A -name=Particion1
		//mkfs -type=full -id=A119
		if continuar {
			break
		}
	}

	fmt.Println("Server started on port 4000")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleRoute)
	router.HandleFunc("/tasks", postMethod).Methods("POST")
	router.HandleFunc("/tasks", getMethod).Methods("GET")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":4000", handler))

}

func AnalyzeQuery(Query string) {
	command, params := getCommandAndParams(Query)
	fmt.Println("comando: ", command, "Parametros: ", params)
	AnalyzeCommnad(command, params)
}

func AnalyzeCommnad(command string, params string) {

	if strings.Contains(command, "mkdisk") {
		fn_mkdisk(params)
	} else if strings.Contains(command, "rmdisk") {
		fn_rmdisk(params)
	} else if strings.Contains(command, "fdisk") {
		fn_fdisk(params)
	} else if strings.Contains(command, "mount") {
		fn_mount(params)
	} else if strings.Contains(command, "unmount") {
		fn_unmount(params)
	} else if strings.Contains(command, "mkfs") {
		fn_mkfs(params)
	} else if strings.Contains(command, "logout") {
		continuar = true
	} else if strings.Contains(command, "#") {
		fn_comentario(params)
	} else if strings.Contains(command, "rep") {
		fn_rep(params)
	} else if strings.Contains(command, "execute") {
		fn_execute("./Analyzer/Entry", "Entry", "mia")
		// modificar funcion fn_execute con los parametros que seria -path=/home/estiben/Documentos/go/src/Proyectos/pruebas.mia
		// fn_execute(params)
	} else if strings.Contains(command, "exit") {
		os.Exit(0)
	} else {
		fmt.Println("Error: comando no encontrado")
	}
}

func fn_mkfs(input string) {
	// Define flags
	fs := flag.NewFlagSet("mkfs", flag.ExitOnError)
	id := fs.String("id", "", "Id")
	type_ := fs.String("type", "", "Tipo")
	fs_ := fs.String("fs", "2fs", "Fs")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := match[2]

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "id", "type", "fs":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
		}
	}

	// Call the function
	FileSystem.Mkfs(*id, *type_, *fs_)

}

func fn_mount(input string) {
	// Define flags
	fs := flag.NewFlagSet("mount", flag.ExitOnError)
	driveletter := fs.String("driveletter", "", "Letra")
	name := fs.String("name", "", "Nombre")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "driveletter", "name":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
		}
	}

	// Call the function
	DiskManagement.Mount(*driveletter, *name)
}

func fn_fdisk(input string) {
	// Define flags
	fs := flag.NewFlagSet("fdisk", flag.ExitOnError)
	size := fs.Int("size", 0, "Tamaño")
	driveletter := fs.String("driveletter", "", "Letra")
	name := fs.String("name", "", "NombredeParticion")
	unit := fs.String("unit", "", "Unidad")
	type_ := fs.String("type", "", "Tipo")
	fit := fs.String("fit", "", "Ajuste")
	delete := fs.String("delete", "", "Eliminar")
	add := fs.Int("add", 0, "Agregar")

	comandoDelete := false
	comandoAdd := false

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])
		flagValue = strings.Trim(flagValue, "\"")
		//fmt.Println("flagName: ", flagName, "flagValue: ", flagValue)
		switch flagName {
		case "size", "fit", "unit", "driveletter", "name", "type", "SIZE", "FIT", "UNIT", "DRIVELETTER", "NAME", "TYPE":
			fs.Set(flagName, flagValue)
		case "delete":
			fs.Set(flagName, flagValue)
			comandoDelete = true
		case "add":
			fs.Set(flagName, flagValue)
			comandoAdd = true
		default:
			fmt.Println("Error: Parametro " + flagName + " no es valido")
			return
		}
	}

	// Validate the flags
	if *size <= 0 && !comandoDelete && !comandoAdd {
		fmt.Println("Parametro size obligatorio revise sintaxis y debe ser mayor a 0")
		return
	}

	// valitate driveletter
	*driveletter = strings.ToUpper(*driveletter)
	if !fileExists(*driveletter) {
		fmt.Println("Error: Driveletter " + *driveletter + " no existe")
		return
	}

	// validate name
	if *name == "" {
		fmt.Println("Error: Name es un parametro obligatorio y no puede estar vacio")
		return
	}
	// validate unit
	*unit = strings.ToUpper(*unit)
	if *unit == "" {
		*unit = "K"
	} else if *unit == "B" {
		// No pasa nada sigue en bytes
	} else if *unit == "K" {
		*size = *size * 1024
	} else if *unit == "M" {
		*size = *size * 1024 * 1024
	} else if *unit != "B" && *unit != "K" && *unit != "M" {
		fmt.Println("Error: Unit debe de ser 'B', 'K', or 'M'")
		return
	}

	// validate type
	*type_ = strings.ToUpper(*type_)
	if *type_ == "" {
		*type_ = "P"
	}
	if *type_ != "P" && *type_ != "E" && *type_ != "L" {
		fmt.Println("Error: Type debe de ser 'P', 'E', o 'L'")
		return
	}

	// validate fit
	*fit = strings.ToUpper(*fit)
	if *fit == "" {
		*fit = "WF"
	} else if *fit != "BF" && *fit != "FF" && *fit != "WF" {
		fmt.Println("Error: Fit debe de ser 'BF', 'FF', o 'WF'")
		return
	}

	// Call the function
	*delete = strings.ToLower(*delete)
	if comandoDelete {
		if *name != "" && *driveletter != "" && *delete == "full" {
			DiskManagement.Fdisk_Delete(*delete, *name, *driveletter)
		} else {
			fmt.Println("Error: Asegurese de que los campos name, driveletter y delete estén llenos y delete sea full")
			return
		}
	} else if comandoAdd {
		if *driveletter != "" && *unit != "" {
			DiskManagement.Fdisk_Add(*add, *unit, *name, *driveletter)
		} else {
			fmt.Println("Error: Solo los campos add y driveletter deben estar llenos")
			return
		}
	} else {
		DiskManagement.Fdisk(*size, *driveletter, *name, *unit, *type_, *fit)
	}

}

// Verifica si un archivo existe
func fileExists(filepath string) bool {
	filepath = strings.ToUpper(filepath)
	filepath = filepath + ".bin"
	path := path.Join("./test/", filepath)
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func fn_mkdisk(params string) {
	// Define flags
	fs := flag.NewFlagSet("mkdisk", flag.ExitOnError)
	size := fs.Int("size", 0, "Tamaño")
	fit := fs.String("fit", "ff", "Ajuste")
	unit := fs.String("unit", "", "Unidad")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(params, -1)

	// Track if we've seen the "size" flag
	sizeSeen := false

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "size", "fit", "unit":
			if flagValue == "" {
				// Aquí puedes poner el código que quieres ejecutar cuando flagValue es una cadena vacía
				fmt.Println("El parametro " + flagName + " no puede estar vacío")
				return
			} else {
				// Si flagValue no está vacío, entonces se establece el valor de la bandera
				sizeSeen = true
				fs.Set(flagName, flagValue)
			}
		default:
			fmt.Println("Error: atributo no encontrado")
			return
		}
	}

	// Check if we've seen the "size" flag
	if !sizeSeen {
		fmt.Println("Error: La bandera 'size' debe aparecer parametro obligatorio")
		return
	} else {
		// Call the function
		DiskManagement.Mkdisk(*size, *fit, *unit)
	}

}

func fn_rmdisk(params string) {
	// Define flags
	fs := flag.NewFlagSet("mkdisk", flag.ExitOnError)
	driveletter := fs.String("driveletter", "", "Letradeunidad")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(params, -1)

	errorRmdisk := false
	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "driveletter":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: atributo no encontrado")
			errorRmdisk = true
			return
		}
	}

	if errorRmdisk {
		return
	} else {
		// Call the function
		DiskManagement.Rmdisk(*driveletter)
	}

}

func fn_unmount(params string) {
	println("unmount", params)
}

// 13 fuctin pause

// 14 function comentario
func fn_comentario(params string) {
	fmt.Println("Comentario: ", params)
}

func fn_execute(path string, filename string, extension string) {
	// Define flags
	fullPath := path + "/" + filename + "." + extension

	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Println("Error al abrir el archivo: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if idx := strings.Index(line, "#"); idx != -1 {
			line = line[:idx]
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// impresion de la linea
		fmt.Println(line)
		// colocar funcion para llamar al analizador
		// analizador(line)
		AnalyzeQuery(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo: ", err)
	}
}

func fn_rep(params string) {
	// Define flags
	fs := flag.NewFlagSet("rep", flag.ExitOnError)
	id := fs.String("id", "", "Id")
	path := fs.String("path", "", "Path")
	name := fs.String("name", "", "Name")
	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(params, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "id", "Path", "name":
			fmt.Println("name: ", name)
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: atributo no encontrado")
			return
		}
	}

	DiskManagement.Rep(*id, *path)
}
