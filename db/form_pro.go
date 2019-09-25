package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// FormPro profesionals form
type FormPro struct {
	refPointer int       `sql:"-"`
	tableName  struct{}  `sql:"form_pro"`
	UserID     int       `sql:"user_id" json:"userId"`
	ID         int       `sql:"id,pk" json:"id"`
	Nro        string    `sql:"nro" json:"nro"`
	Pro1       string    `sql:"pro1" json:"pro1"`
	Pro2       string    `sql:"pro2" json:"pro2"`
	Pro3       string    `sql:"pro3" json:"pro3"`
	Pro4       string    `sql:"pro4" json:"pro4"`
	Pro5       string    `sql:"pro5" json:"pro5"`
	Pro6       string    `sql:"pro6" json:"pro6"`
	Pro7       string    `sql:"pro7" json:"pro7"`
	Pro8       string    `sql:"pro8" json:"pro8"`
	Pro9       string    `sql:"pro9" json:"pro9"`
	Pro10      string    `sql:"pro10" json:"pro10"`
	Pro11      string    `sql:"pro11" json:"pro11"`
	Pro12      string    `sql:"pro12" json:"pro12"`
	Pro13      string    `sql:"pro13" json:"pro13"`
	Pro14      string    `sql:"pro14" json:"pro14"`
	Pro15      string    `sql:"pro15" json:"pro15"`
	Pro16      string    `sql:"pro16" json:"pro16"`
	Pro17      string    `sql:"pro17" json:"pro17"`
	Pro18      string    `sql:"pro18" json:"pro18"`
	Pro19      string    `sql:"pro19" json:"pro19"`
	Pro20      string    `sql:"pro20" json:"pro20"`
	Pro21      string    `sql:"pro21" json:"pro21"`
	Pro22      string    `sql:"pro22" json:"pro22"`
	Pro23      string    `sql:"pro23" json:"pro23"`
	Pro24      string    `sql:"pro24" json:"pro24"`
	Pro25      string    `sql:"pro25" json:"pro25"`
	Pro26      string    `sql:"pro26" json:"pro26"`
	Pro27      string    `sql:"pro27" json:"pro27"`
	Pro28      string    `sql:"pro28" json:"pro28"`
	Pro29      string    `sql:"pro29" json:"pro29"`
	Pro30      string    `sql:"pro30" json:"pro30"`
	Pro31      string    `sql:"pro31" json:"pro31"`
	Pro32      string    `sql:"pro32" json:"pro32"`
	Pro33      string    `sql:"pro33" json:"pro33"`
	Pro34      string    `sql:"pro34" json:"pro34"`
	Pro35      string    `sql:"pro35" json:"pro35"`
	Pro36      string    `sql:"pro36" json:"pro36"`
	Pro37      string    `sql:"pro37" json:"pro37"`
	Pro38      string    `sql:"pro38" json:"pro38"`
	Pro39      string    `sql:"pro39" json:"pro39"`
	Pro40      string    `sql:"pro40" json:"pro40"`
	Pro41      string    `sql:"pro41" json:"pro41"`
	Pro42      string    `sql:"pro42" json:"pro42"`
	Pro43      string    `sql:"pro43" json:"pro43"`
	Pro44      string    `sql:"pro44" json:"pro44"`
	Pro45      string    `sql:"pro45" json:"pro45"`
	Pro46      string    `sql:"pro46" json:"pro46"`
	Pro47      string    `sql:"pro47" json:"pro47"`
	Pro48      string    `sql:"pro48" json:"pro48"`
	Pro49      string    `sql:"pro49" json:"pro49"`
	Pro50      string    `sql:"pro50" json:"pro50"`
	Pro51      string    `sql:"pro51" json:"pro51"`
	Pro52      string    `sql:"pro52" json:"pro52"`
	Pro53      string    `sql:"pro53" json:"pro53"`
	Pro54      string    `sql:"pro54" json:"pro54"`
	Pro55      string    `sql:"pro55" json:"pro55"`
	Pro56      string    `sql:"pro56" json:"pro56"`
	Pro57      string    `sql:"pro57" json:"pro57"`
	Pro58      string    `sql:"pro58" json:"pro58"`
	Pro59      string    `sql:"pro59" json:"pro59"`
	CreatedAt  time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at" json:"updated_at"`
}

// Save saves the profesional form
func (fpro *FormPro) Save() error {
	err := DBCon.Insert(fpro)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}

// CreateFormProTable creates the table
func CreateFormProTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&FormPro{}, opts)
	if err != nil {
		log.Panicf("error while creating table database %v\n", err)
		return err
	}
	log.Printf("Table FormPro created")
	return nil
}

// GetByNro get form by nro
func (fpro *FormPro) GetByNro(nro string) error {
	err := DBCon.Model(fpro).Where("nro = ?", nro).First()
	if err != nil {
		return err
	}
	return nil
}

// Update updates the profesiona form
func (fpro *FormPro) Update() error {
	var formProAux FormPro
	formProAux.ID = fpro.ID
	if err := formProAux.GetByID(); err != nil {
		return err
	}
	fpro.UserID = formProAux.UserID
	fpro.CreatedAt = formProAux.CreatedAt
	fpro.UpdatedAt = time.Now()
	if err := DBCon.Update(fpro); err != nil {
		return err
	}
	return nil
}

// Delete deletes the profesional form
func (fpro *FormPro) Delete() error {
	if err := DBCon.Delete(fpro); err != nil {
		return err
	}
	return nil
}

// GetByID get form by id
func (fpro *FormPro) GetByID() error {
	if err := DBCon.Select(fpro); err != nil {
		return err
	}
	return nil
}

// GetQuestions get the form pro questions
func (fpro *FormPro) GetQuestions() (interface{}, error) {
	type oneStruct struct {
		Pro7  Select       `json:"pro7"`
		Pro8  Select       `json:"pro8"`
		Pro9  SelectOption `json:"pro9"`
		Pro10 Select       `json:"pro10"`
		Pro16 SelectOption `json:"pro16"`
	}
	type twoStruct struct {
		Pro18 Select       `json:"pro18"`
		Pro21 SelectOption `json:"pro21"`
		Pro22 Select       `json:"pro22"`
		Pro23 Select       `json:"pro23"`
		Pro24 Select       `json:"pro24"`
		Pro25 SelectOption `json:"pro25"`
		Pro26 SelectOption `json:"pro26"`
		Pro27 SelectOption `json:"pro27"`
		Pro28 SelectOption `json:"pro28"`
		Pro29 SelectOption `json:"pro29"`
		Pro30 SelectOption `json:"pro30"`
		Pro31 Multiple     `json:"pro31"`
	}
	type threeStruct struct {
		Pro32 Multiple `json:"pro32"`
	}
	type fourStruct struct {
		Pro34 Multiple `json:"pro34"`
	}
	type fiveStruct struct {
		Pro39 Multiple     `json:"pro39"`
		Pro41 SelectOption `json:"pro41"`
		Pro42 SelectOption `json:"pro42"`
		Pro43 SelectOption `json:"pro43"`
	}
	type sixStruct struct {
		Pro44 SelectOption `json:"pro44"`
		Pro45 SelectOption `json:"pro45"`
		Pro46 SelectOption `json:"pro46"`
		Pro47 SelectOption `json:"pro47"`
		Pro48 SelectOption `json:"pro48"`
	}
	type sevenStruct struct {
		Pro49 Multiple     `json:"pro49"`
		Pro50 Multiple     `json:"pro50"`
		Pro51 SelectOption `json:"pro51"`
	}
	type eightStruct struct {
		Pro52 SelectOption `json:"pro52"`
	}

	one := oneStruct{}
	one.Pro7.Title = "Género"
	one.Pro8.Title = "Estado civil"
	one.Pro9.Title = "Universidad donde se tituló"
	one.Pro10.Title = "Rango de edad"
	one.Pro16.Title = "Certificaciones con las que cuenta."
	two := twoStruct{}
	two.Pro18.Title = "Actualmente, trabaja:"
	two.Pro21.Title = "Prestaciones laborales"
	two.Pro22.Title = "Relación de su trabajo con el área de Informática"
	two.Pro23.Title = "Salario mensual en Bolivianos"
	two.Pro24.Title = "Alcance de la institución"
	two.Pro25.Title = "Tipos de empleo"
	two.Pro26.Title = "Rubro o actividad de la institución"
	two.Pro27.Title = "Cargo que ocupa"
	two.Pro28.Title = "Forma de ingreso a la institución"
	two.Pro29.Title = "Tipo de institución"
	two.Pro30.Title = "Actividades que desarrolla en su trabajo"
	two.Pro31.Title = "Aspectos de su formación académica que son útiles en su desempeño laboral"
	three := threeStruct{}
	three.Pro32.Title = "Perfil profesional en el mercado laboral."
	four := fourStruct{}
	four.Pro34.Title = "Plan de estudios."
	five := fiveStruct{}
	five.Pro39.Title = "Respecto al mercado laboral"
	five.Pro41.Title = "¿Qué denominaciones considera que son y serán requeridas por el mercado laboral?"
	five.Pro42.Title = "La especialización en el área de Informática debería hacerse:"
	five.Pro43.Title = "¿Con cuál de las siguientes modalidades se tituló?"
	six := sixStruct{}
	six.Pro44.Title = "Lenguajes de programación que utiliza en su trabajo"
	six.Pro45.Title = "Gestores de base de datos que utiliza en su trabajo"
	six.Pro46.Title = "Metodologías de desarrollo de sistemas que utiliza en su trabajo"
	six.Pro47.Title = "Sistemas operativos que utiliza en su trabajo"
	six.Pro48.Title = "Frameworks que utiliza en su trabajo"
	seven := sevenStruct{}
	seven.Pro49.Title = "Competencias generales que se busca en el mercado laboral profesional"
	seven.Pro50.Title = "Competencias específicas que se busca en el mercado laboral profesional"
	seven.Pro51.Title = "¿Cuáles son las áreas y tecnologías que tienen y tendrán mayor impacto en el mercado laboral?"
	eight := eightStruct{}
	eight.Pro52.Title = "El los siguientes 5 años te gustaría trabajar en:"

	if err := one.Pro7.parseSimple("pro7", fpro); err != nil {
		return nil, err
	}
	if err := one.Pro8.parseSimple("pro8", fpro); err != nil {
		return nil, err
	}
	if err := one.Pro9.parse("pro9", fpro); err != nil {
		return nil, err
	}
	if err := one.Pro10.parseSimple("pro10", fpro); err != nil {
		return nil, err
	}
	if err := one.Pro16.parseMul("pro16", fpro); err != nil {
		return nil, err
	}

	if err := two.Pro18.parseSimple("pro18", fpro); err != nil {
		return nil, err
	}
	if err := two.Pro21.parseMul("pro21", fpro); err != nil {
		return nil, err
	}
	if err := two.Pro22.parseSimple("pro22", fpro); err != nil {
		return nil, err
	}
	if err := two.Pro23.parseSimple("pro23", fpro); err != nil {
		return nil, err
	}
	if err := two.Pro24.parseSimple("pro24", fpro); err != nil {
		return nil, err
	}
	if err := two.Pro25.parse("pro25", fpro); err != nil {
		return nil, err
	}
	if err := two.Pro26.parse("pro26", fpro); err != nil {
		return nil, err
	}
	if err := two.Pro27.parseMul("pro27", fpro); err != nil {
		return nil, err
	}
	if err := two.Pro28.parse("pro28", fpro); err != nil {
		return nil, err
	}
	if err := two.Pro29.parse("pro29", fpro); err != nil {
		return nil, err
	}
	if err := two.Pro30.parseMul("pro30", fpro); err != nil {
		return nil, err
	}
	if err := two.Pro31.parse("pro31", fpro, false); err != nil {
		return nil, err
	}

	if err := three.Pro32.parse("pro32", fpro, false); err != nil {
		return nil, err
	}

	if err := four.Pro34.parse("pro34", fpro, false); err != nil {
		return nil, err
	}

	if err := five.Pro39.parse("pro39", fpro, false); err != nil {
		return nil, err
	}
	if err := five.Pro41.parseMul("pro41", fpro); err != nil {
		return nil, err
	}
	if err := five.Pro42.parse("pro42", fpro); err != nil {
		return nil, err
	}
	if err := five.Pro43.parse("pro43", fpro); err != nil {
		return nil, err
	}

	if err := six.Pro44.parseMul("pro44", fpro); err != nil {
		return nil, err
	}
	if err := six.Pro45.parseMul("pro45", fpro); err != nil {
		return nil, err
	}
	if err := six.Pro46.parseMul("pro46", fpro); err != nil {
		return nil, err
	}
	if err := six.Pro47.parseMul("pro47", fpro); err != nil {
		return nil, err
	}
	if err := six.Pro48.parseMul("pro48", fpro); err != nil {
		return nil, err
	}

	if err := seven.Pro49.parse("pro49", fpro, true); err != nil {
		return nil, err
	}
	if err := seven.Pro50.parse("pro50", fpro, true); err != nil {
		return nil, err
	}
	if err := seven.Pro51.parseMul("pro51", fpro); err != nil {
		return nil, err
	}

	if err := eight.Pro52.parseMul("pro52", fpro); err != nil {
		return nil, err
	}

	resp := []RespStruct{}
	resp = append(resp, RespStruct{"Aspectos Generales", one})
	resp = append(resp, RespStruct{"Situación Laboral", two})
	resp = append(resp, RespStruct{"Perfil Profesional", three})
	resp = append(resp, RespStruct{"Plan de Estudios con el que se Formó", four})
	resp = append(resp, RespStruct{"Grados, Titulación y Menciones", five})
	resp = append(resp, RespStruct{"Áreas de Conocimiento", six})
	resp = append(resp, RespStruct{"Competencias del Profesional en el Área de Informática y Sistemas", seven})
	resp = append(resp, RespStruct{"Aspiraciones Futuras", eight})

	return resp, nil
}
