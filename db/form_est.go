package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// FormEst handles student form
type FormEst struct {
	refPointer int       `sql:"-"`
	tableName  struct{}  `sql:"form_est" pg:",discard_unknown_columns"`
	UserID     int       `sql:"user_id" json:"userId"`
	ID         int       `sql:"id,pk" json:"id"`
	Nro        string    `sql:"nro" json:"nro"`
	Est1       string    `sql:"est1" json:"est1"`
	Est2       string    `sql:"est2" json:"est2"`
	Est3       string    `sql:"est3" json:"est3"`
	Est4       string    `sql:"est4" json:"est4"`
	Est5       string    `sql:"est5" json:"est5"`
	Est6       string    `sql:"est6" json:"est6"`
	Est7       string    `sql:"est7" json:"est7"`
	Est8       string    `sql:"est8" json:"est8"`
	Est9       string    `sql:"est9" json:"est9"`
	Est10      string    `sql:"est10" json:"est10"`
	Est11      string    `sql:"est11" json:"est11"`
	Est12      string    `sql:"est12" json:"est12"`
	Est13      string    `sql:"est13" json:"est13"`
	Est14      string    `sql:"est14" json:"est14"`
	Est15      string    `sql:"est15" json:"est15"`
	Est16      string    `sql:"est16" json:"est16"`
	Est17      string    `sql:"est17" json:"est17"`
	Est18      string    `sql:"est18" json:"est18"`
	Est19      string    `sql:"est19" json:"est19"`
	Est20      string    `sql:"est20" json:"est20"`
	Est21      string    `sql:"est21" json:"est21"`
	Est22      string    `sql:"est22" json:"est22"`
	Est23      string    `sql:"est23" json:"est23"`
	Est24      string    `sql:"est24" json:"est24"`
	Est25      string    `sql:"est25" json:"est25"`
	Est26      string    `sql:"est26" json:"est26"`
	Est27      string    `sql:"est27" json:"est27"`
	Est28      string    `sql:"est28" json:"est28"`
	Est29      string    `sql:"est29" json:"est29"`
	Est30      string    `sql:"est30" json:"est30"`
	Est31      string    `sql:"est31" json:"est31"`
	Est32      string    `sql:"est32" json:"est32"`
	Est33      string    `sql:"est33" json:"est33"`
	Est34      string    `sql:"est34" json:"est34"`
	Est35      string    `sql:"est35" json:"est35"`
	Est36      string    `sql:"est36" json:"est36"`
	Est37      string    `sql:"est37" json:"est37"`
	Est38      string    `sql:"est38" json:"est38"`
	Est39      string    `sql:"est39" json:"est39"`
	Est40      string    `sql:"est40" json:"est40"`
	Est41      string    `sql:"est41" json:"est41"`
	Est42      string    `sql:"est42" json:"est42"`
	Est43      string    `sql:"est43" json:"est43"`
	Est44      string    `sql:"est44" json:"est44"`
	Est45      string    `sql:"est45" json:"est45"`
	Est46      string    `sql:"est46" json:"est46"`
	Est47      string    `sql:"est47" json:"est47"`
	Est48      string    `sql:"est48" json:"est48"`
	Est49      string    `sql:"est49" json:"est49"`
	Est50      string    `sql:"est50" json:"est50"`
	Est51      string    `sql:"est51" json:"est51"`
	Est52      string    `sql:"est52" json:"est52"`
	Est53      string    `sql:"est53" json:"est53"`
	Est54      string    `sql:"est54" json:"est54"`
	Est55      string    `sql:"est55" json:"est55"`
	Est56      string    `sql:"est56" json:"est56"`
	Est57      string    `sql:"est57" json:"est57"`
	Est58      string    `sql:"est58" json:"est58"`
	Est59      string    `sql:"est59" json:"est59"`
	Est60      string    `sql:"est60" json:"est60"`
	Est61      string    `sql:"est61" json:"est61"`
	Est62      string    `sql:"est62" json:"est62"`
	Est63      string    `sql:"est63" json:"est63"`
	Est64      string    `sql:"est64" json:"est64"`
	Est65      string    `sql:"est65" json:"est65"`
	Est66      string    `sql:"est66" json:"est66"`
	Est67      string    `sql:"est67" json:"est67"`
	Est68      string    `sql:"est68" json:"est68"`
	CreatedAt  time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at" json:"updated_at"`
}

// Save saves the form
func (fe *FormEst) Save(db *pg.DB) error {
	err := db.Insert(fe)
	if err != nil {
		log.Printf("Error")
		return err
	}
	return nil
}

// Update updates the form
func (fe *FormEst) Update() error {
	var estAux FormEst
	estAux.ID = fe.ID
	if err := estAux.GetByID(); err != nil {
		return err
	}
	fe.UserID = estAux.UserID
	fe.CreatedAt = estAux.CreatedAt
	fe.UpdatedAt = time.Now()
	err := DBCon.Update(fe)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes the form
func (fe *FormEst) Delete() error {
	err := DBCon.Delete(fe)
	if err != nil {
		return err
	}
	return nil
}

// CreateFormEstTable creates the table
func CreateFormEstTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&FormEst{}, opts)
	if err != nil {
		log.Panicf("error while creating table database %v\n", err)
		return err
	}
	log.Printf("Table FormEst created")
	return nil
}

// GetByNro get form by nro
func (fe *FormEst) GetByNro(nro string) error {
	err := DBCon.Model(fe).Where("nro = ?", nro).First()
	if err != nil {
		return err
	}
	return nil
}

// GetByID get form by id
func (fe *FormEst) GetByID() error {
	if err := DBCon.Select(fe); err != nil {
		return err
	}
	return nil
}

// GetQuestions get the form est questions all in one
func (fe *FormEst) GetQuestions() (interface{}, error) {
	// RESPONSES STRUCTS
	type oneStruct struct {
		Est4  Select       `json:"est4"`
		Est10 Select       `json:"est10"`
		Est11 Select       `json:"est11"`
		Est13 SelectOption `json:"est13"`
		Est14 SelectOption `json:"est14"`
		Est15 Select       `json:"est15"`
	}
	type twoStruct struct {
		Est16 Select       `json:"est16"`
		Est18 SelectOption `json:"est18"`
		Est19 SelectOption `json:"est19"`
		Est21 SelectOption `json:"est21"`
	}
	type respStruct struct {
		Name string      `json:"name"`
		Data interface{} `json:"data"`
	}
	// TITLE DEFINITIONS STRUCTS
	one := oneStruct{}
	one.Est4.Title = "Número de estudiantes por semestre a la gestión."
	one.Est10.Title = "Cantidad de estudiantes varones y cantidad de estudiantes mujeres."
	one.Est11.Title = "Cantidad de estudiantes procedentes de unidades educativas fiscales, particulares y de convenio."
	one.Est13.Title = "Residencia de los estudiantes."
	one.Est14.Title = "Tipos de acceso a la red Internet."
	one.Est15.Title = "Cantidad de estudiantes solteros, casados, divorciados y viudos."
	two := twoStruct{}
	two.Est16.Title = "Cantidad de estudiantes que ingresaron por curso pre-universitario, examen de dispensación, adminsión directa, traspaso o paralela."
	two.Est18.Title = "Top de áreas de preferencia de los estudiantes."
	two.Est19.Title = "Principales razones por las que los estudiantes abandonan materias."
	two.Est21.Title = "Principales razones por las que los estudiantes abandonan la carrera."
	if err := one.Est4.parseSimple("est4", fe); err != nil {
		return nil, err
	}
	if err := one.Est10.parseSimple("est10", fe); err != nil {
		return nil, err
	}
	if err := one.Est11.parseSimple("est11", fe); err != nil {
		return nil, err
	}
	if err := one.Est13.parse("est13", fe); err != nil {
		return nil, err
	}
	if err := one.Est14.parse("est14", fe); err != nil {
		return nil, err
	}
	if err := one.Est15.parseSimple("est15", fe); err != nil {
		return nil, err
	}

	if err := two.Est16.parseSimple("est16", fe); err != nil {
		return nil, err
	}
	if err := two.Est18.parseMul("est18", fe); err != nil {
		return nil, err
	}
	if err := two.Est19.parseMul("est19", fe); err != nil {
		return nil, err
	}
	if err := two.Est21.parseMul("est21", fe); err != nil {
		return nil, err
	}
	// FINAL RESPONSE STRUCT BUILD
	resp := []respStruct{}
	resp = append(resp, respStruct{"Aspectos Generales", one})
	resp = append(resp, respStruct{"Vinculación con la Carrera", two})
	return resp, nil
}
