package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// FormPre formulario de los prefas
type FormPre struct {
	refPointer int       `sql:"-"`
	tableName  struct{}  `sql:"form_pre"`
	UserID     int       `sql:"user_id" json:"userId"`
	ID         int       `sql:"id,pk" json:"id"`
	Nro        string    `sql:"nro" json:"nro"`
	Pre1       string    `sql:"pre1" json:"pre1"`
	Pre2       string    `sql:"pre2" json:"pre2"`
	Pre3       string    `sql:"pre3" json:"pre3"`
	Pre4       string    `sql:"pre4" json:"pre4"`
	Pre5       string    `sql:"pre5" json:"pre5"`
	Pre6       string    `sql:"pre6" json:"pre6"`
	Pre7       string    `sql:"pre7" json:"pre7"`
	Pre8       string    `sql:"pre8" json:"pre8"`
	Pre9       string    `sql:"pre9" json:"pre9"`
	Pre10      string    `sql:"pre10" json:"pre10"`
	Pre11      string    `sql:"pre11" json:"pre11"`
	Pre12      string    `sql:"pre12" json:"pre12"`
	Pre13      string    `sql:"pre13" json:"pre13"`
	Pre14      string    `sql:"pre14" json:"pre14"`
	Pre15      string    `sql:"pre15" json:"pre15"`
	Pre16      string    `sql:"pre16" json:"pre16"`
	Pre17      string    `sql:"pre17" json:"pre17"`
	Pre18      string    `sql:"pre18" json:"pre18"`
	Pre19      string    `sql:"pre19" json:"pre19"`
	Pre20      string    `sql:"pre20" json:"pre20"`
	Pre21      string    `sql:"pre21" json:"pre21"`
	Pre22      string    `sql:"pre22" json:"pre22"`
	Pre23      string    `sql:"pre23" json:"pre23"`
	Pre24      string    `sql:"pre24" json:"pre24"`
	Pre25      string    `sql:"pre25" json:"pre25"`
	Pre26      string    `sql:"pre26" json:"pre26"`
	Pre26a     string    `sql:"pre26a" json:"pre26a"`
	Pre27      string    `sql:"pre27" json:"pre27"`
	Pre27a     string    `sql:"pre27a" json:"pre27a"`
	Pre28      string    `sql:"pre28" json:"pre28"`
	Pre29      string    `sql:"pre29" json:"pre29"`
	Pre30      string    `sql:"pre30" json:"pre30"`
	Pre31      string    `sql:"pre31" json:"pre31"`
	Pre32      string    `sql:"pre32" json:"pre32"`
	Pre33      string    `sql:"pre33" json:"pre33"`
	Pre34      string    `sql:"pre34" json:"pre34"`
	Pre35      string    `sql:"pre35" json:"pre35"`
	Pre36      string    `sql:"pre36" json:"pre36"`
	Pre37      string    `sql:"pre37" json:"pre37"`
	Pre38      string    `sql:"pre38" json:"pre38"`
	Pre39      string    `sql:"pre39" json:"pre39"`
	Pre40      string    `sql:"pre40" json:"pre40"`
	Pre41      string    `sql:"pre41" json:"pre41"`
	Pre42      string    `sql:"pre42" json:"pre42"`
	CreatedAt  time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt  time.Time `sql:"updated_at" json:"updated_at"`
}

// CreateFormPreTable creates form_pre table
func CreateFormPreTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&FormPre{}, opts)
	if err != nil {
		log.Panicf("error while creating table database %v\n", err)
		return err
	}
	log.Printf("Table FormPre created")
	return nil
}

// Save saves a form in database
func (fpre *FormPre) Save() error {
	err := DBCon.Insert(fpre)
	if err != nil {
		return err
	}
	return nil
}

// GetByNro get form by nro
func (fpre *FormPre) GetByNro(nro string) error {
	err := DBCon.Model(fpre).Where("nro = ?", nro).First()
	if err != nil {
		return err
	}
	return nil
}

// Update updates the prefas form
func (fpre *FormPre) Update() error {
	var formPreAux FormPre
	if err := formPreAux.GetByNro(fpre.Nro); err != nil {
		return err
	}
	fpre.UserID = formPreAux.UserID
	fpre.CreatedAt = formPreAux.CreatedAt
	fpre.UpdatedAt = time.Now()
	if err := DBCon.Update(fpre); err != nil {
		return err
	}
	return nil
}

// Delete deletes the prefas form
func (fpre *FormPre) Delete() error {
	if err := DBCon.Delete(fpre); err != nil {
		return err
	}
	return nil
}

// GetByID get form by ID
func (fpre *FormPre) GetByID() error {
	if err := DBCon.Select(fpre); err != nil {
		return err
	}
	return nil
}

// GetQuestions get all questions
func (fpre *FormPre) GetQuestions() (interface{}, error) {
	type oneStruct struct {
		Pre9  Select       `json:"pre9"`
		Pre10 SelectOption `json:"pre10"`
		Pre11 SelectOption `json:"pre11"`
		Pre12 SelectOption `json:"pre12"`
		Pre13 SelectOption `json:"pre13"`
		Pre14 Select       `json:"pre14"`
	}
	type twoStruct struct {
		Pre16 SelectOption `json:"pre16"`
		Pre17 SelectOption `json:"pre17"`
		Pre18 SelectOption `json:"pre18"`
		Pre25 Multiple     `json:"pre25"`
		Pre26 Select       `json:"pre26"`
		Pre27 Select       `json:"pre27"`
	}
	type threeStruct struct {
		Pre33 SelectOption `json:"pre33"`
		Pre34 Select       `json:"pre34"`
		Pre35 SelectOption `json:"pre35"`
		Pre36 SelectOption `json:"pre36"`
	}
	type fourStruct struct {
		Pre37 Multiple     `json:"pre37"`
		Pre38 SelectOption `json:"pre38"`
	}
	type fiveStruct struct {
		Pre39 SelectOption `json:"pre39"`
	}
	one := oneStruct{}
	one.Pre9.Title = "Género"
	one.Pre10.Title = "Tipo unidad educativa"
	one.Pre11.Title = "Ubicación unidad educativa"
	one.Pre12.Title = "Lugar de residencia"
	one.Pre13.Title = "Desde su casa accede a Internet por:"
	one.Pre14.Title = "Estado civil"
	two := twoStruct{}
	two.Pre16.Title = "Áreas en las que le gustaría profundizar"
	two.Pre17.Title = "Medio por el que se informó sobre la carrera de Informática"
	two.Pre18.Title = "Participación en concursos de programación"
	two.Pre25.Title = "Importancia de las asignaturas para la carrera de Informática"
	two.Pre26.Title = "¿Cuántas veces ha cursado el curso pre-universitario?"
	two.Pre27.Title = "¿Cuántas veces se ha presentado al examen de dispensación?"
	three := threeStruct{}
	three.Pre33.Title = "Actualmente"
	three.Pre34.Title = "Relación de su trabajo con la carrera"
	three.Pre35.Title = "Razones por las que trabaja"
	three.Pre36.Title = "Tipo de empleo"
	four := fourStruct{}
	four.Pre37.Title = "Licenciado en Informática"
	four.Pre38.Title = "Con qué denominaciones de la carrera de Informática esta de acuerdo"
	five := fiveStruct{}
	five.Pre39.Title = "¿Cuándo termine la carrera, me gustaría trabajar en?"
	if err := one.Pre9.parseSimple("pre9", fpre); err != nil {
		return nil, err
	}
	if err := one.Pre10.parse("pre10", fpre); err != nil {
		return nil, err
	}
	if err := one.Pre11.parse("pre11", fpre); err != nil {
		return nil, err
	}
	if err := one.Pre12.parse("pre12", fpre); err != nil {
		return nil, err
	}
	if err := one.Pre13.parse("pre13", fpre); err != nil {
		return nil, err
	}
	if err := one.Pre14.parseSimple("pre14", fpre); err != nil {
		return nil, err
	}
	if err := two.Pre16.parseMul("pre16", fpre); err != nil {
		return nil, err
	}
	if err := two.Pre17.parseMul("pre17", fpre); err != nil {
		return nil, err
	}
	if err := two.Pre18.parseMul("pre18", fpre); err != nil {
		return nil, err
	}
	if err := two.Pre25.parse("pre25", fpre, true); err != nil {
		return nil, err
	}
	if err := two.Pre26.parseSimple("pre26", fpre); err != nil {
		return nil, err
	}
	if err := two.Pre27.parseSimple("pre27", fpre); err != nil {
		return nil, err
	}
	if err := three.Pre33.parse("pre33", fpre); err != nil {
		return nil, err
	}
	if err := three.Pre34.parseSimple("pre34", fpre); err != nil {
		return nil, err
	}
	if err := three.Pre35.parse("pre35", fpre); err != nil {
		return nil, err
	}
	if err := three.Pre36.parse("pre36", fpre); err != nil {
		return nil, err
	}
	if err := four.Pre37.parse("pre37", fpre, false); err != nil {
		return nil, err
	}
	if err := four.Pre38.parseMul("pre38", fpre); err != nil {
		return nil, err
	}
	if err := five.Pre39.parseMul("pre39", fpre); err != nil {
		return nil, err
	}
	resp := []RespStruct{}
	resp = append(resp, RespStruct{"Aspectos Generales", one})
	resp = append(resp, RespStruct{"Vinculación con la Carrera", two})
	resp = append(resp, RespStruct{"Sitación Laboral", three})
	resp = append(resp, RespStruct{"Grados, Titulación y Menciones", four})
	resp = append(resp, RespStruct{"Aspiraciones Futuras", five})
	return resp, nil
}
