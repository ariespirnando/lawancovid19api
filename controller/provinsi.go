package controller
import ( 
	"net/http" 
	"log" 
	"github.com/ariespirnando/covid19/utils" 
	"github.com/ariespirnando/covid19/model" 
	"github.com/ariespirnando/covid19/config"
	"github.com/gin-gonic/gin" 
)

func Allprovinsi(c *gin.Context){   
	var data model.RespProvinsi
	var arrData []model.RespProvinsi
	DB := config.Connect() 
	defer DB.Close()  
	rows, err := DB.Query("call `lihat_provinsi`")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{  
			"error_code" : "000002",    
		})
	}else{
		for rows.Next() { 
			if err2 := rows.Scan(&data.KodeProvinsi,&data.NamaProvinsi); err != nil {
				log.Fatal(err2.Error()) 
			} else { 
				arrData = append(arrData,data)
			}
		}  
		c.JSON(http.StatusOK, gin.H{  
			"error_code" : "000000",   
			"data": arrData,
		})
	} 
				
}
func Cronjob(c *gin.Context){  
	utils.RunUpdateDatatoDB() 
	c.JSON(http.StatusOK, gin.H{  
		"error_code" : "000000",   
	}) 		
}

 


