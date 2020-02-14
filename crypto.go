package IBank_core
import(
	"crypto/md5"
	"encoding/hex"
)

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
//////XML
//data, err = xml.Marshal(ATMs)
//if err != nil {
//log.Fatal(err)
//return err
//}
//err = ioutil.WriteFile("ATM.xml", data, 0666)
//if err != nil {
//log.Fatal(err)
//return err
//}
//return nil
//}
//
//func AddAtmFromXmlJson(db *sql.DB) (err error) {
//	/////XML
//	file, err := ioutil.ReadFile("ATM.xml")
//	if err != nil {
//		log.Fatalf("Can't read file %e", err)
//		return err
//	}
//	var Atms cmodels.AtmList
//	err = xml.Unmarshal(file, &Atms)
//	if err != nil {
//		log.Fatal("Can't Unmarshal file", err)
//		return err
//	}
//	for _, Atm := range Atms.ATMs{
//		Address := Atm.Name
//		Locked := Atm.Locked
//		err = dbupdate.AddATM(Address, Locked, db)
//		if err != nil {
//			log.Printf("Проблема соединения с сервером %e", err)
//			return err
//		}
//	}