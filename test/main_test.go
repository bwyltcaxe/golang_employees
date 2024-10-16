package person_test

import (
    "testing"
    "github.com/bwyltcaxe/golang_employees/internal/person"
)

func TestPersonSetNameSuccess(test *testing.T) {
    obj := person.Person{Name: "Ivan",
        Gender: "man",
        ID:     0x1080}

    names := []string{"Egor", "Timur", "Matvey"}

    for i, v := range names {
        obj.SetName(v)
        if obj.Name != v {
            test.Errorf("[%d]: Expected %s instead of %s", i, v, obj.Name)
        }
    }
}

func TestPersonSetNameFail(test *testing.T) {
    obj := person.Person{Name: "Ivan",
        Gender: "man",
        ID:     0x1080}

    names := []string{"Egor", "Timur", "Matvey"}

    for i, v := range names {
        obj.SetName(v)
        // Let's inverse the previous results.
        if obj.Name == v {
            test.Errorf("[%d]: Expected %s instead of %s", i, v, obj.Name)
        }
    }
}

func TestPersonSetId(test *testing.T) {
    obj := person.Person{Name: "Ivan",
        Gender: "man",
        ID:     0x1080}

    names := []int{0x1080, 0x8016, 0x8016}

    for i, v := range names {
        obj.SetID(v)
        if obj.ID != v {
            test.Errorf("[%d]: Expected %d instead of %d", i, v, obj.ID)
        }
    }
}

func TestPersonSetGender(test *testing.T) {
    obj := person.Person{Name: "Ivan",
        Gender: "man",
        ID:     0x1080}

    names := []string{"man", "woman", "woman"}

    for i, v := range names {
        obj.SetGender(v)
        if obj.Gender != v {
            test.Errorf("[%d]: Expected %s instead of %s", i, v, obj.Gender)
        }
    }
}
