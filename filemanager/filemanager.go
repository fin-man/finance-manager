package filemanager
 
import(
	"os"
)

type FileManager struct {
}

func (f *FileManager) OpenFile(filename string, flags int, permissions os.FileMode)(*os.File,error){
	return os.OpenFile(filename,flags , permissions)
} 
