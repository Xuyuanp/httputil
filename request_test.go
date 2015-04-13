package httputil

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"sync"
	"testing"
)

type Person struct {
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}

var p = Person{Name: "Lilei", Age: 15}

func TestParseJSON(t *testing.T) {
	data, _ := json.Marshal(p)
	buffer := bytes.NewBuffer(data)
	req, _ := http.NewRequest("POST", "http://localhost", buffer)

	np := &Person{}
	if err := ParseJSON(req, np); err != nil {
		t.Error("parse json failed:", err)
	}
	if np.Name != p.Name || np.Age != p.Age {
		t.Error("parse json result doesn't match")
	}
}

func TestParseXML(t *testing.T) {
	data, _ := xml.Marshal(p)
	buffer := bytes.NewBuffer(data)
	req, _ := http.NewRequest("POST", "http://localhost", buffer)

	np := &Person{}
	if err := ParseXML(req, np); err != nil {
		t.Error("parse xml failed:", err)
	}
	if np.Name != p.Name || np.Age != p.Age {
		t.Error("parse xml result doesn't match")
	}
}

func TestPostJSON(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			t.Errorf("wrong method. POST expected, %s got", req.Method)
		}
		ct := req.Header.Get("Content-Type")
		if ct != "application/json; charset=utf-8" {
			t.Errorf("wrong header(Content-Type). 'application/json; charset=utf-8' expected, '%s' got", ct)
		}
		np := &Person{}
		if err := ParseJSON(req, np); err != nil {
			t.Error("parse json failed:", err)
		}
		if np.Name != p.Name || np.Age != p.Age {
			t.Errorf("wrong json content. %+v expected, %+v got", p, np)
		}
		wg.Done()
	})
	go http.ListenAndServe(":12345", nil)

	if _, err := PostJSON("http://127.0.0.1:12345", p); err != nil {
		t.Error(err)
	}
	wg.Wait()
}
