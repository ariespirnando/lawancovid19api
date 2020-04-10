package utils
import (
	"os"
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"time"
	"github.com/ariespirnando/covid19/config" 
	"github.com/ariespirnando/covid19/model" 
)

func UpdateDataNegara(){
	url := os.Getenv("URL_GATEWAY") 
	
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println("Http Request Error : ",err)
	}else{
		data, _ := ioutil.ReadAll(resp.Body)
		jsonData := string(data)  
		var Explode []model.Negara
		var err = json.Unmarshal([]byte(jsonData),&Explode)
		if err != nil{
			fmt.Println(err.Error())
		} 
		var SaveData model.SaveorUpdatedata
		SaveData.KodeProvinsi = 0
		SaveData.NamaProvinsi = Explode[0].Nama
		SaveData.KasusPositif = ConvertToInt(Explode[0].Positif)
		SaveData.KasusMeninggal = ConvertToInt(Explode[0].Meninggal)
		SaveData.KasusSehat = ConvertToInt(Explode[0].Sembuh)
		go UpdateDB(&SaveData)
	}
	defer resp.Body.Close() 
}
func UpdateDataProvinsi(){
	resp, err := http.Get(os.Getenv("URL_GATEWAY")+"/provinsi")
	if err != nil{
		fmt.Println("Http Request Error : ",err)
	}else{
		data, _ := ioutil.ReadAll(resp.Body)
		jsonData := string(data)  
		var Explode []model.Provinsi
		var err = json.Unmarshal([]byte(jsonData),&Explode)
		if err != nil{
			fmt.Println(err.Error())
		}
		for i, _ := range Explode { 
			var SaveData model.SaveorUpdatedata
			SaveData.KodeProvinsi = Explode[i].Attributes.KodeProvi
			SaveData.NamaProvinsi = Explode[i].Attributes.Provinsi
			SaveData.KasusPositif = Explode[i].Attributes.KasusPosi
			SaveData.KasusMeninggal = Explode[i].Attributes.KasusMeni
			SaveData.KasusSehat = Explode[i].Attributes.KasusSemb
			if(i%2==0){
				time.Sleep(11 * time.Second)
			}
			UpdateDB(&SaveData)
		}  
	}
	defer resp.Body.Close()
}


func UpdateDB(data *model.SaveorUpdatedata){
	DB := config.Connect()
    //defer DB.Close() 
	_, err := DB.Query("CALL `update_data`(?,?,?,?,?)",
							data.KodeProvinsi,
							data.NamaProvinsi,
							data.KasusPositif,
							data.KasusSehat,
							data.KasusMeninggal, 
						)
	if err != nil {
		log.Print(err.Error())
	} 
}

func RunUpdateDatatoDB(){
	go UpdateDataNegara()
	go UpdateDataProvinsi()
}


