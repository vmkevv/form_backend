package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// FormDoc doc form structure
type FormDoc struct {
	refPointer int       `sql:"-"`
	tableName  struct{}  `sql:"form_doc"`
	UserID     int       `sql:"user_id" json:"userId"`
	ID         int       `sql:"id,pk" json:"id"`
	Nro        string    `sql:"nro" json:"nro"`
	Doc1       string    `sql:"doc1" json:"doc1"`
	Doc2       string    `sql:"doc2" json:"doc2"`
	Doc3       string    `sql:"doc3" json:"doc3"`
	Doc4       string    `sql:"doc4" json:"doc4"`
	Doc5       string    `sql:"doc5" json:"doc5"`
	Doc6       string    `sql:"doc6" json:"doc6"`
	Doc7       string    `sql:"doc7" json:"doc7"`
	Doc8       string    `sql:"doc8" json:"doc8"`
	Doc9       string    `sql:"doc9" json:"doc9"`
	Doc10      string    `sql:"doc10" json:"doc10"`
	Doc11      string    `sql:"doc11" json:"doc11"`
	Doc12      string    `sql:"doc12" json:"doc12"`
	Doc13      string    `sql:"doc13" json:"doc13"`
	Doc14a     string    `sql:"doc14a" json:"doc14a"`
	Doc14b     string    `sql:"doc14b" json:"doc14b"`
	Doc14c     string    `sql:"doc14c" json:"doc14c"`
	Doc14d     string    `sql:"doc14d" json:"doc14d"`
	Doc15      string    `sql:"doc15" json:"doc15"`
	Doc16      string    `sql:"doc16" json:"doc16"`
	Doc17      string    `sql:"doc17" json:"doc17"`
	Doc18      string    `sql:"doc18" json:"doc18"`
	Doc19      string    `sql:"doc19" json:"doc19"`
	Doc20      string    `sql:"doc20" json:"doc20"`
	Doc21      string    `sql:"doc21" json:"doc21"`
	Doc22      string    `sql:"doc22" json:"doc22"`
	Doc23      string    `sql:"doc23" json:"doc23"`
	Doc24      string    `sql:"doc24" json:"doc24"`
	Doc24a     string    `sql:"doc24a" json:"doc24a"`
	Doc25      string    `sql:"doc25" json:"doc25"`
	Doc25a     string    `sql:"doc25a" json:"doc25a"`
	Doc26      string    `sql:"doc26" json:"doc26"`
	Doc26a     string    `sql:"doc26a" json:"doc26a"`
	Doc27      string    `sql:"doc27" json:"doc27"`
	Doc27a     string    `sql:"doc27a" json:"doc27a"`
	Doc28      string    `sql:"doc28" json:"doc28"`
	Doc29      string    `sql:"doc29" json:"doc29"`
	Doc30      string    `sql:"doc30" json:"doc30"`
	Doc31      string    `sql:"doc31" json:"doc31"`
	Doc32      string    `sql:"doc32" json:"doc32"`
	Doc33      string    `sql:"doc33" json:"doc33"`
	Doc34      string    `sql:"doc34" json:"doc34"`
	Doc35      string    `sql:"doc35" json:"doc35"`
	Doc36      string    `sql:"doc36" json:"doc36"`
	Doc37      string    `sql:"doc37" json:"doc37"`
	Doc38      string    `sql:"doc38" json:"doc38"`
	Doc39      string    `sql:"doc39" json:"doc39"`
	Doc40      string    `sql:"doc40" json:"doc40"`
	Doc41      string    `sql:"doc41" json:"doc41"`
	Doc42      string    `sql:"doc42" json:"doc42"`
	Doc43      string    `sql:"doc43" json:"doc43"`
	Doc44      string    `sql:"doc44" json:"doc44"`
	Doc45      string    `sql:"doc45" json:"doc45"`
	Doc46      string    `sql:"doc46" json:"doc46"`
	Doc47      string    `sql:"doc47" json:"doc47"`
	CreatedAt  time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at" json:"updated_at"`
}

// CreateFormDocTable creates doc table
func CreateFormDocTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	if err := db.CreateTable(&FormDoc{}, opts); err != nil {
		log.Panicf("error while creating table database %v\n", err)
		return err
	}
	log.Printf("Table FormDoc created")
	return nil
}

// Save saves the doc form
func (fdoc *FormDoc) Save() error {
	err := DBCon.Insert(fdoc)
	if err != nil {
		return err
	}
	return nil
}

// GetByNro get doc form by nro
func (fdoc *FormDoc) GetByNro(nro string) error {
	err := DBCon.Model(fdoc).Where("nro = ?", nro).First()
	if err != nil {
		return err
	}
	return nil
}

// Update updater the doc form
func (fdoc *FormDoc) Update() error {
	var forDocAux FormDoc
	if err := forDocAux.GetByNro(fdoc.Nro); err != nil {
		return err
	}
	fdoc.UserID = forDocAux.UserID
	fdoc.CreatedAt = forDocAux.CreatedAt
	fdoc.UpdatedAt = time.Now()
	if err := DBCon.Update(fdoc); err != nil {
		return err
	}
	return nil
}

// Delete deletes the doc form
func (fdoc *FormDoc) Delete() error {
	if err := DBCon.Delete(fdoc); err != nil {
		return err
	}
	return nil
}

// GetByID gets for by Id
func (fdoc *FormDoc) GetByID() error {
	if err := DBCon.Select(fdoc); err != nil {
		return err
	}
	return nil
}

// GetQuestions get all questions
func (fdoc *FormDoc) GetQuestions() (interface{}, error) {
	type oneStruct struct {
		Doc5 Select       `json:"doc5"`
		Doc6 SelectOption `json:"doc6"`
		Doc7 SelectOption `json:"doc7"`
		Doc8 Select       `json:"doc8"`
		Doc9 SelectOption `json:"doc9"`
	}
	type twoStruct struct {
		Doc16 SelectOption `json:"doc16"`
		Doc17 SelectOption `json:"doc17"`
	}
	type threeStruct struct {
		Doc18 Multiple `json:"doc18"`
		Doc19 Multiple `json:"doc19"`
	}
	type fourStruct struct {
		Doc21 Multiple `json:"doc21"`
	}
	type fiveStruct struct {
		Doc28 Multiple `json:"doc28"`
	}
	type sixStruct struct {
		Doc30 Multiple     `json:"doc30"`
		Doc32 Multiple     `json:"doc32"`
		Doc33 SelectOption `json:"doc33"`
		Doc34 SelectOption `json:"doc34"`
		Doc36 SelectOption `json:"doc36"`
	}
	type sevenStruct struct {
		Doc37 SelectOption `json:"doc37"`
		Doc38 SelectOption `json:"doc38"`
		Doc39 SelectOption `json:"doc39"`
		Doc40 SelectOption `json:"doc40"`
		Doc41 SelectOption `json:"doc41"`
	}
	type eightStruct struct {
		Doc42 Multiple     `json:"doc42"`
		Doc43 Multiple     `json:"doc43"`
		Doc44 SelectOption `json:"doc44"`
	}
	one := oneStruct{}
	one.Doc5.Title = "Género"
	one.Doc6.Title = "Residencia actual"
	one.Doc7.Title = "Universidad dónde se tituló"
	one.Doc8.Title = "Procedencia"
	one.Doc9.Title = "Profesiones"
	two := twoStruct{}
	two.Doc16.Title = "Categoría docente"
	two.Doc17.Title = "Áreas en las que se desempeña actualmente"
	three := threeStruct{}
	three.Doc18.Title = "Capacidades del licenciado en Informática"
	three.Doc19.Title = "Licenciado en informática y la realidad actual"
	four := fourStruct{}
	four.Doc21.Title = "Características del plan de estudios actual"
	five := fiveStruct{}
	five.Doc28.Title = "Características infraestructura"
	six := sixStruct{}
	six.Doc30.Title = "Menciones y mercado laboral"
	six.Doc32.Title = "Denominaciones de carreras en el área de Informática"
	six.Doc33.Title = "La especialización en el área de Informática debería hacerse:"
	six.Doc34.Title = "Modalidades de gracuación adecuadas para la carrera"
	six.Doc36.Title = "Cuánto demora un estudiante en realizar su trabajo de grado"
	seven := sevenStruct{}
	seven.Doc37.Title = "En cuanto a lenguages de programación"
	seven.Doc38.Title = "En cuanto a gestores de bases de datos"
	seven.Doc39.Title = "En cuanto a metodologías de desarrollo de sistemas"
	seven.Doc40.Title = "En cuanto a sistemas operativos"
	seven.Doc41.Title = "En cuanto a Frameworks"
	eight := eightStruct{}
	eight.Doc42.Title = "Competencias generales que debe tener el profesional en informática"
	eight.Doc43.Title = "Competencias específicas que debe tener el profesional en informática"
	eight.Doc44.Title = "Líneas de investigación"
	if err := one.Doc5.parseSimple("doc5", fdoc); err != nil {
		return nil, err
	}
	if err := one.Doc6.parse("doc6", fdoc); err != nil {
		return nil, err
	}
	if err := one.Doc7.parseMul("doc7", fdoc); err != nil {
		return nil, err
	}
	if err := one.Doc8.parseSimple("doc8", fdoc); err != nil {
		return nil, err
	}
	if err := one.Doc9.parseMul("doc9", fdoc); err != nil {
		return nil, err
	}
	if err := two.Doc16.parseMul("doc16", fdoc); err != nil {
		return nil, err
	}
	if err := two.Doc17.parseMul("doc17", fdoc); err != nil {
		return nil, err
	}
	if err := three.Doc18.parse("doc18", fdoc, false); err != nil {
		return nil, err
	}
	if err := three.Doc19.parse("doc19", fdoc, false); err != nil {
		return nil, err
	}
	if err := four.Doc21.parse("doc21", fdoc, false); err != nil {
		return nil, err
	}
	if err := five.Doc28.parse("doc28", fdoc, false); err != nil {
		return nil, err
	}
	if err := six.Doc30.parse("doc30", fdoc, false); err != nil {
		return nil, err
	}
	if err := six.Doc32.parse("doc32", fdoc, false); err != nil {
		return nil, err
	}
	if err := six.Doc33.parse("doc33", fdoc); err != nil {
		return nil, err
	}
	if err := six.Doc34.parseMul("doc34", fdoc); err != nil {
		return nil, err
	}
	if err := six.Doc36.parse("doc36", fdoc); err != nil {
		return nil, err
	}
	if err := seven.Doc37.parseMul("doc37", fdoc); err != nil {
		return nil, err
	}
	if err := seven.Doc38.parseMul("doc38", fdoc); err != nil {
		return nil, err
	}
	if err := seven.Doc39.parseMul("doc39", fdoc); err != nil {
		return nil, err
	}
	if err := seven.Doc40.parseMul("doc40", fdoc); err != nil {
		return nil, err
	}
	if err := seven.Doc41.parseMul("doc41", fdoc); err != nil {
		return nil, err
	}
	if err := eight.Doc42.parse("doc42", fdoc, true); err != nil {
		return nil, err
	}
	if err := eight.Doc43.parse("doc43", fdoc, true); err != nil {
		return nil, err
	}
	if err := eight.Doc44.parseMul("doc44", fdoc); err != nil {
		return nil, err
	}

	resp := []RespStruct{}
	resp = append(resp, RespStruct{"Aspectos Generales", one})
	resp = append(resp, RespStruct{"Vinculación con la Carrera", two})
	resp = append(resp, RespStruct{"Perfil Profesional", three})
	resp = append(resp, RespStruct{"Plan de Estudio Actual", four})
	resp = append(resp, RespStruct{"Infraestructura (Aulas, Laboratorios, Biblioteca, Servicios)", five})
	resp = append(resp, RespStruct{"Grados, Titulación y Menciones", six})
	resp = append(resp, RespStruct{"Áreas de Conocimiento", seven})
	resp = append(resp, RespStruct{"Competencias de Profesional en el Área de la Informática y Sistemas", eight})
	return resp, nil
}
