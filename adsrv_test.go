package adsrv

import (
	"testing"
)

func TestSortPriority(t *testing.T) {
	d := new("", "", SRVSet{
		{Target: "a.example.org", Port: 389, Priority: 9, Weight: 100},
		{Target: "e.example.org", Port: 389, Priority: 5, Weight: 100},
		{Target: "c.example.org", Port: 389, Priority: 7, Weight: 100},
		{Target: "b.example.org", Port: 389, Priority: 8, Weight: 100},
		{Target: "d.example.org", Port: 389, Priority: 6, Weight: 100},
		{Target: "f.example.org", Port: 389, Priority: 4, Weight: 100},
	})
	if d.SRV[0].Target != "f.example.org" {
		t.Logf("Expected f.example.org, got %s", d.SRV[0].Target)
		t.Fail()
	}
}

func TestSortWeight(t *testing.T) {
	d := new("", "", SRVSet{
		{Target: "a.example.org", Port: 389, Priority: 0, Weight: 100},
		{Target: "e.example.org", Port: 389, Priority: 0, Weight: 500},
		{Target: "c.example.org", Port: 389, Priority: 0, Weight: 300},
		{Target: "f.example.org", Port: 389, Priority: 0, Weight: 600},
		{Target: "b.example.org", Port: 389, Priority: 0, Weight: 200},
		{Target: "d.example.org", Port: 389, Priority: 0, Weight: 400},
	})
	if d.SRV[0].Target != "f.example.org" {
		t.Logf("Expected f.example.org, got %s", d.SRV[0].Target)
		t.Fail()
	}
}

func TestDialFail(t *testing.T) {
	d := new("", "", SRVSet{
		{Target: "a.example.org", Port: 389, Priority: 0, Weight: 100},
		{Target: "e.example.org", Port: 389, Priority: 0, Weight: 500},
		{Target: "c.example.org", Port: 389, Priority: 0, Weight: 300},
		{Target: "f.example.org", Port: 389, Priority: 0, Weight: 600},
		{Target: "b.example.org", Port: 389, Priority: 0, Weight: 200},
		{Target: "d.example.org", Port: 389, Priority: 0, Weight: 400},
	})
	if _, err := d.Dial(); err == nil {
		t.Logf("Expected Dial() to fail")
		t.Fail()
	}
}

func TestMakeSiteDomain(t *testing.T) {
	got := makeSiteDomain("win.slac.stanford.edu", "SLAC")
	if got != "SLAC._sites.win.slac.stanford.edu" {
		t.Logf("Expected SLAC._sites.win.slac.stanford.edu, got %v", got)
		t.Fail()
	}
}
