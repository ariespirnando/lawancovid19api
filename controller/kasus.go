package controller
import ( 
	"net/http" 
	"log"  
	"github.com/ariespirnando/covid19/model" 
	"github.com/ariespirnando/covid19/config"
	"github.com/gin-gonic/gin" 
)

func Kasusperprovinsi(c *gin.Context){   
	var json model.ReqKasusProvinsi  
    if err := c.ShouldBindJSON(&json);err != nil{ 
        c.JSON(http.StatusBadRequest, gin.H{  
            "error_code" : "000001", 
        }) 
    }else{
		var data model.Respkasusprovinsi 
		var grafik model.RespGrafik
		var arr_grafik []model.RespGrafik
		DB := config.Connect() 
		defer DB.Close()   
		err := DB.QueryRow("call `lihat_kasus_perprovinsi`(?)",json.KodeProvinsi, 
				).Scan(&data.KasusPositif,
					&data.KasusSembuh, 
					&data.KasusMeninggal,
					&data.KodeProvinsi,
					&data.NamaProvinsi)
		rows, err2 := DB.Query("call `lihat_grafik_perprov`(?)",json.KodeProvinsi)   
		if err != nil || err2 != nil {
			log.Print(err)
			c.JSON(http.StatusMethodNotAllowed, gin.H{  
				"error_code" : "000002",    
			})
		}else{

			for rows.Next() { 
				if err3 := rows.Scan(&grafik.Tanggal,&grafik.KasusPositif); err3 != nil {
					log.Fatal(err3.Error()) 
				} else { 
					arr_grafik = append(arr_grafik,grafik)
				}
			}
			data.Grafik = arr_grafik
			c.JSON(http.StatusOK, gin.H{  
				"error_code" : "000000",   
				"data": data,
			}) 
		} 
	}		
}

func Kasusprovinsi(c *gin.Context){   
	var data model.Respkasusprovinsi  
	var datalist []model.Respkasusprovinsi  

	DB := config.Connect() 
	defer DB.Close() 
	rows, err := DB.Query("call `list_kasus_all`")   
	if err != nil{
		log.Print(err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{  
			"error_code" : "000002",    
		})
	}else{
		for rows.Next() { 
			if err1 := rows.Scan(&data.KodeProvinsi,
				&data.NamaProvinsi,
				&data.KasusPositif,
				&data.KasusSembuh, 
				&data.KasusMeninggal,); err1 != nil {
				log.Fatal(err1.Error()) 
			} else { 
				datalist = append(datalist,data)
			}
		} 
		c.JSON(http.StatusOK, gin.H{  
			"error_code" : "000000",   
			"data": datalist,
		}) 
	}
}