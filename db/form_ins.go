package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// FormIns institutions form
type FormIns struct {
	refPointer int       `sql:"-"`
	tableName  struct{}  `sql:"form_ins"`
	UserID     int       `sql:"user_id" json:"userId"`
	ID         int       `sql:"id,pk" json:"id"`
	Nro        string    `sql:"nro" json:"nro"`
	Ins1       string    `sql:"ins1" json:"ins1"`
	Ins2       string    `sql:"ins2" json:"ins2"`
	Ins3       string    `sql:"ins3" json:"ins3"`
	Ins4       string    `sql:"ins4" json:"ins4"`
	Ins5       string    `sql:"ins5" json:"ins5"`
	Ins6       string    `sql:"ins6" json:"ins6"`
	Ins7       string    `sql:"ins7" json:"ins7"`
	Ins8       string    `sql:"ins8" json:"ins8"`
	Ins9       string    `sql:"ins9" json:"ins9"`
	Ins10      string    `sql:"ins10" json:"ins10"`
	Ins11      string    `sql:"ins11" json:"ins11"`
	Ins12      string    `sql:"ins12" json:"ins12"`
	Ins13a     string    `sql:"ins13a" json:"ins13a"`
	Ins13b     string    `sql:"ins13b" json:"ins13b"`
	Ins13c     string    `sql:"ins13c" json:"ins13c"`
	Ins13d     string    `sql:"ins13d" json:"ins13d"`
	Ins13e     string    `sql:"ins13e" json:"ins13e"`
	Ins13f     string    `sql:"ins13f" json:"ins13f"`
	Ins13g     string    `sql:"ins13g" json:"ins13g"`
	Ins13h     string    `sql:"ins13h" json:"ins13h"`
	Ins13i     string    `sql:"ins13i" json:"ins13i"`
	Ins13j     string    `sql:"ins13j" json:"ins13j"`
	Ins13k     string    `sql:"ins13k" json:"ins13k"`
	Ins13l     string    `sql:"ins13l" json:"ins13l"`
	Ins13m     string    `sql:"ins13m" json:"ins13m"`
	Ins14      string    `sql:"ins14" json:"ins14"`
	Ins15      string    `sql:"ins15" json:"ins15"`
	Ins15a     string    `sql:"ins15a" json:"ins15a"`
	Ins16      string    `sql:"ins16" json:"ins16"`
	Ins17      string    `sql:"ins17" json:"ins17"`
	Ins18      string    `sql:"ins18" json:"ins18"`
	Ins19      string    `sql:"ins19" json:"ins19"`
	Ins20      string    `sql:"ins20" json:"ins20"`
	Ins21      string    `sql:"ins21" json:"ins21"`
	Ins22      string    `sql:"ins22" json:"ins22"`
	Ins22a     string    `sql:"ins22a" json:"ins22a"`
	Ins23      string    `sql:"ins23" json:"ins23"`
	Ins23a     string    `sql:"ins23a" json:"ins23a"`
	Ins24      string    `sql:"ins24" json:"ins24"`
	Ins25      string    `sql:"ins25" json:"ins25"`
	Ins26      string    `sql:"ins26" json:"ins26"`
	Ins27      string    `sql:"ins27" json:"ins27"`
	Ins28      string    `sql:"ins28" json:"ins28"`
	Ins29      string    `sql:"ins29" json:"ins29"`
	Ins30      string    `sql:"ins30" json:"ins30"`
	Ins31      string    `sql:"ins31" json:"ins31"`
	Ins32      string    `sql:"ins32" json:"ins32"`
	Ins33      string    `sql:"ins33" json:"ins33"`
	Ins34      string    `sql:"ins34" json:"ins34"`
	Ins35      string    `sql:"ins35" json:"ins35"`
	Ins36      string    `sql:"ins36" json:"ins36"`
	Ins37      string    `sql:"ins37" json:"ins37"`
	CreatedAt  time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at" json:"updated_at"`
}

// Save saves the institution form
func (fins *FormIns) Save() error {
	if err := DBCon.Insert(fins); err != nil {
		return err
	}
	return nil
}

// CreateFormInsTable creates the table
func CreateFormInsTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&FormIns{}, opts)
	if err != nil {
		log.Panicf("error while creating table database %v\n", err)
		return err
	}
	log.Printf("Table FormIns created")
	return nil
}

// Update updates the form
func (fins *FormIns) Update() error {
	var formInsAux FormIns
	formInsAux.ID = fins.ID
	if err := formInsAux.GetByID(); err != nil {
		return err
	}
	fins.UserID = formInsAux.UserID
	fins.CreatedAt = formInsAux.CreatedAt
	fins.UpdatedAt = time.Now()
	err := DBCon.Update(fins)
	if err != nil {
		return err
	}
	return nil
}

// GetByNro get form by nro
func (fins *FormIns) GetByNro(nro string) error {
	err := DBCon.Model(fins).Where("nro = ?", nro).First()
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes the form
func (fins *FormIns) Delete() error {
	err := DBCon.Delete(fins)
	if err != nil {
		return err
	}
	return nil
}

// GetByID get form by ID
func (fins *FormIns) GetByID() error {
	if err := DBCon.Select(fins); err != nil {
		return err
	}
	return nil
}

// GetQuestions get form ins question reports
func (fins *FormIns) GetQuestions() (interface{}, error) {
	type oneStruct struct {
		Ins8  SelectOption `json:"ins8"`
		Ins9  SelectOption `json:"ins9"`
		Ins10 Select       `json:"ins10"`
		Ins11 Select       `json:"ins11"`
	}
	type twoStruct struct {
		Ins14 SelectOption `json:"ins14"`
		Ins16 SelectOption `json:"ins16"`
	}
	type threeStruct struct {
		Ins17 Multiple `json:"ins17"`
	}
	type fourStruct struct {
		Ins19 Multiple `json:"ins19"`
	}
	type fiveStruct struct {
		Ins24 Multiple     `json:"ins24"`
		Ins26 SelectOption `json:"ins26"`
	}
	type sixStruct struct {
		Ins27 SelectOption `json:"ins27"`
		Ins28 SelectOption `json:"ins28"`
		Ins29 SelectOption `json:"ins29"`
		Ins30 SelectOption `json:"ins30"`
		Ins31 SelectOption `json:"ins31"`
	}
	type sevenStruct struct {
		Ins32 Multiple     `json:"ins32"`
		Ins33 Multiple     `json:"ins33"`
		Ins34 SelectOption `json:"ins34"`
	}

	one := oneStruct{}
	one.Ins8.Title = "Tipo de institución"
	one.Ins9.Title = "Rubro de la institución"
	one.Ins10.Title = "Nro de empleados en la institución"
	one.Ins11.Title = "Alcance de la institución"
	two := twoStruct{}
	two.Ins14.Title = "Forma en que se accede a los cargos"
	two.Ins16.Title = "Principales actividades que desarrolla el personal del área de Informática y Sistemas en la institución"
	three := threeStruct{}
	three.Ins17.Title = "Mercado laboral"
	four := fourStruct{}
	four.Ins19.Title = "Planes de Estudio"
	five := fiveStruct{}
	five.Ins24.Title = "Valoración"
	five.Ins26.Title = "¿Qué denominaciones considera que son y serán requeridas por el mercado laboral?"
	six := sixStruct{}
	six.Ins27.Title = "Lenguages de programación que utiliza la institución"
	six.Ins28.Title = "Gestores de base de datos que utiliza la institución"
	six.Ins29.Title = "Metodologías de desarrollo de sistemas que utiliza la institución"
	six.Ins30.Title = "Sistemas Operativos que utiliza la institución"
	six.Ins31.Title = "Frameworks que utiliza la institución"
	seven := sevenStruct{}
	seven.Ins32.Title = "Competencias generales que busca la institución"
	seven.Ins33.Title = "Competencias específicas que busca la institución"
	seven.Ins34.Title = "Áreas y tecnologías que tienen y tendrán mayor impacto en el mercado laboral"

	if err := one.Ins8.parse("ins8", fins); err != nil {
		return nil, err
	}
	if err := one.Ins9.parseMul("ins9", fins); err != nil {
		return nil, err
	}
	if err := one.Ins10.parseSimple("ins10", fins); err != nil {
		return nil, err
	}
	if err := one.Ins11.parseSimple("ins11", fins); err != nil {
		return nil, err
	}

	if err := two.Ins14.parseMul("ins14", fins); err != nil {
		return nil, err
	}
	if err := two.Ins16.parseMul("ins16", fins); err != nil {
		return nil, err
	}

	if err := three.Ins17.parse("ins17", fins, false); err != nil {
		return nil, err
	}

	if err := four.Ins19.parse("ins19", fins, false); err != nil {
		return nil, err
	}

	if err := five.Ins24.parse("ins24", fins, false); err != nil {
		return nil, err
	}
	if err := five.Ins26.parseMul("ins24", fins); err != nil {
		return nil, err
	}

	if err := six.Ins27.parseMul("ins27", fins); err != nil {
		return nil, err
	}
	if err := six.Ins28.parseMul("ins28", fins); err != nil {
		return nil, err
	}
	if err := six.Ins29.parseMul("ins29", fins); err != nil {
		return nil, err
	}
	if err := six.Ins30.parseMul("ins30", fins); err != nil {
		return nil, err
	}
	if err := six.Ins31.parseMul("ins31", fins); err != nil {
		return nil, err
	}

	if err := seven.Ins32.parse("ins32", fins, true); err != nil {
		return nil, err
	}
	if err := seven.Ins33.parse("ins33", fins, true); err != nil {
		return nil, err
	}
	if err := seven.Ins34.parseMul("ins34", fins); err != nil {
		return nil, err
	}

	resp := []RespStruct{}
	resp = append(resp, RespStruct{"Aspectos Generales", one})
	resp = append(resp, RespStruct{"Aspectos Laborales", two})
	resp = append(resp, RespStruct{"Perfil Profesional", three})
	resp = append(resp, RespStruct{"Planes de Estudio", four})
	resp = append(resp, RespStruct{"Grados, Titulación y Menciones", five})
	resp = append(resp, RespStruct{"Áreas de Conocimiento", six})
	resp = append(resp, RespStruct{"Competencias del Profesional en el Área de Informática y Sistemas", seven})
	return resp, nil
}
