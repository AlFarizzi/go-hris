package helper

import (
	"net/http"
	"text/template"
)

func AuthViewParser(w http.ResponseWriter, view string) {
	tmp := template.Must(template.ParseGlob("template/auth/*.html"))
	tmp.ExecuteTemplate(w, view, nil)
}

func KaryawanViewParser(w http.ResponseWriter, view string, data map[string]interface{}) {
	tmp := template.Must(template.ParseGlob("template/karyawan/*.html"))
	tmp = template.Must(tmp.ParseGlob("template/dashboard_partials/*.html"))
	tmp.ExecuteTemplate(w, view, data)
}

func PositionViewParser(w http.ResponseWriter, view string, data map[string]interface{}) {
	tmp := template.Must(template.ParseGlob("template/job_position/*.html"))
	tmp = template.Must(tmp.ParseGlob("template/dashboard_partials/*.html"))
	tmp.ExecuteTemplate(w, view, data)
}

func HubunganKeluargaViewParser(w http.ResponseWriter, view string, data map[string]interface{}) {
	tmp := template.Must(template.ParseGlob("template/hubungan_keluarga/*.html"))
	tmp = template.Must(tmp.ParseGlob("template/dashboard_partials/*.html"))
	tmp.ExecuteTemplate(w, view, data)
}

func JenisKelaminViewParser(w http.ResponseWriter, view string, data map[string]interface{}) {
	tmp := template.Must(template.ParseGlob("template/jenis_kelamin/*.html"))
	tmp = template.Must(tmp.ParseGlob("template/dashboard_partials/*.html"))
	tmp.ExecuteTemplate(w, view, data)
}
