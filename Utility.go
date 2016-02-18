package ChronosDB
import (                                                                        
//        "fmt" 
//        "log" 
        "time" 
//        "strconv" 
) 

func getDateStartSec(timestamp int64) ( int64, int64 ) { 
    const shortForm = "2006-01-02"                                             
        tm := time.Unix(timestamp/1000, 0) 
        DateStart :=  tm.Format(shortForm)
        st, _ := time.Parse(shortForm, DateStart) 
        dateZeroOclock := st.UTC().Unix() * 1000 
        //fmt.Println(st.UTC().Unix()) 
        //fmt.Println(timestamp) 
        //fmt.Println(timestamp - 1000 * st.UTC().Unix() ) 
        return dateZeroOclock, timestamp - dateZeroOclock 
} 
