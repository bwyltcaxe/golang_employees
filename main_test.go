package main

import (
    "testing"
)

func TestPersonSetNameSuccess(test *testing.T) {
    obj := Person{Name: "Ivan",
        Gender: man,
        Id:     0x1080}

    names := []string{"Egor", "Timur", "Matvey"}

    for i, v := range names {
        obj.SetName(v)
        if obj.Name != v {
            test.Errorf("[%d]: Expected %s instead of %s", i, v, obj.Name)
        }
    }
}

func TestPersonSetNameFail(test *testing.T) {
    obj := Person{Name: "Ivan",
        Gender: man,
        Id:     0x1080}

    names := []string{"Egor", "Timur", "Matvey"}

    for i, v := range names {
        obj.SetName(v)
        /* Let's inverse the previous results.  */
        if obj.Name == v {
            test.Errorf("[%d]: Expected %s instead of %s", i, v, obj.Name)
        }
    }
}

func TestPersonSetId(test *testing.T) {
    obj := Person{Name: "Ivan",
        Gender: man,
        Id:     0x1080}

    names := []int{0x1080, 0x8016, 0x8016}

    for i, v := range names {
        obj.SetId(v)
        if obj.Id != v {
            test.Errorf("[%d]: Expected %d instead of %d", i, v, obj.Id)
        }
    }
}

func TestPersonSetGender(test *testing.T) {
    obj := Person{Name: "Ivan",
        Gender: man,
        Id:     0x1080}

    names := []GenderEnum{man, woman, woman}

    for i, v := range names {
        obj.SetGender(v)
        if obj.Gender != v {
            test.Errorf("[%d]: Expected %d instead of %d", i, v, obj.Gender)
        }
    }
}
