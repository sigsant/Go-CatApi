package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	facts "getURL/api"
)

var (
	cantidad = flag.Int("c", 1, "Cantidad de hechos a mostrar")
)

var (
	errNotArg    = errors.New("No se ha introducido ningún argumento")
	errParseJSON = errors.New("ERROR: Se ha producido un error en la conversión del JSON: ")
	errEndPoint  = errors.New("ERROR: No se ha podido acceder a la información de la API. Verifica el endpoint")
)

func concatenar(argumento ...string) (string, error) {
	if len(argumento) == 0 {
		return "", errNotArg
	}
	return strings.Join(argumento, " "), nil
}

func hasTimeOut(err error) bool {
	//Compara el error con su tipo
	switch err := err.(type) {
	case *url.Error: // En caso de error generico al parsear un html o query
		if err, ok := err.Err.(net.Error); ok && err.Timeout() { // ¿Error de conectividad en red y en tiempo de espera?
			return true
		}
	case net.Error:
		if err.Timeout() { // ¿Error en conexion?
			return true
		}
	case *net.OpError:
		if err.Timeout() { //¿Error de acceso a la DNS?
			return true
		}
	}
	errTxt := " use of closed network connection"

	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}

func main() {

	flag.Parse()

	amount := strconv.Itoa(*cantidad)

	catFactSg := facts.Facts{}
	catFactPl := []facts.Facts{} //Convierte el struct en una lista de structs
	testURL := "https://cat-fact.herokuapp.com/facts/random?animal_type=cat&amount=" + amount

	response, err := http.Get(testURL)
	if err != nil && hasTimeOut(err) {
		fmt.Printf("\n\tERROR: No puede conectarse a %s : %s", testURL, response.Status)
		fmt.Println("\n\t", hasTimeOut(err))
		os.Exit(1)
	}
	defer response.Body.Close()
	fmt.Printf("\n\tINFO: Conexión a %s: %s\n", testURL, response.Status)

	//TODO! Si el endpoint de la API es erronea aún parsea el HTML de la pagina incorrecta
	data, err := ioutil.ReadAll(response.Body)
	if err != nil || bytes.Contains(data, []byte("error-pages")) {
		fmt.Println("\n\t", errEndPoint)
		os.Exit(1)
	}
	// fmt.Println("\n\n", string(data[:]))

	if *cantidad < 2 {
		_ = json.Unmarshal(data, &catFactSg)
		fmt.Println("\n\n\tID: ", catFactSg.ID)
		fmt.Println("\n\tFact: ", catFactSg.Fact)
	}

	_ = json.Unmarshal(data, &catFactPl)

	for i := range catFactPl {
		fmt.Println("\n\n\tID: ", catFactPl[i].ID)
		fmt.Println("\n\tFact: ", catFactPl[i].Fact)
	}
}
