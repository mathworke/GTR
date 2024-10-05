package assets

import (
	"encoding/xml"
	"io"
	"log"
	"os"
)

const (
	LOG_PATH = "logs"  // путь/до/папки/логов
	LOG_FILE = "/.log" // /название_файла.log
)

type Logger struct {
	Log *os.File

	logging bool

	logInfo  *log.Logger
	logError *log.Logger
}

func NewLogger(logging bool) *Logger {
	os.Mkdir(LOG_PATH, 066)
	file, err := os.OpenFile(LOG_PATH+LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return &Logger{
		Log:      file,
		logging:  logging,
		logInfo:  log.New(file, "[INFO]\t", log.Ldate|log.Ltime),
		logError: log.New(file, "[ERROR]\t", log.Ldate|log.Ltime),
	}
}

func (l *Logger) LogPrint(format string, args ...any) {
	l.logInfo.Printf(" "+format, args...)
}

func (l *Logger) LogIngo(format string, args ...any) {
	if l.logging {
		l.logInfo.Printf(" "+format, args...)
	}
}

func (l *Logger) LogError(format string, args ...any) {
	l.logError.Printf(" "+format, args...)
}

func (l *Logger) PANIC(format string, args ...any) {
	l.logError.Panicf(" "+format, args...)
	os.Exit(1)
}

// struct for config file
type Modules struct {
	XMLName xml.Name `xml:"modules"`
	Modules []struct {
		XMLName     xml.Name `xml:"module"`
		Type        string   `xml:"type,attr"`
		Description string   `xml:"description"`
		Tests       struct {
			XMLName  xml.Name `xml:"tests"`
			TestCase []struct {
				XMLName xml.Name `xml:"test_case"`
				Id      string   `xml:"id"`
				Title   string   `xml:"title"`
			} `xml:"test_case"`
		} `xml:tests`
	} `xml:"module"`
}

type Person struct {
	XMLName xml.Name `xml:"person"`
	Fio     string   `xml:"fio"`
	Post    string   `xml:"post"`
}

func (m *Modules) GetConfiguration(path string, logger *Logger) {
	getDataFromXml(path, m, logger)
}

func (p *Person) GetPerson(path string, logger *Logger) {
	getDataFromXml(path, p, logger)
}

func getDataFromXml(path string, source any, logger *Logger) {
	xmlFile, err := os.Open(path)
	if err != nil {
		logger.PANIC(err.Error())
	}
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	err = xml.Unmarshal(byteValue, &source)
	if err != nil {
		logger.PANIC(err.Error())
	}
}
