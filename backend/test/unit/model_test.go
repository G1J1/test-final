package unit

import (
	"testing"

	
	"github.com/G1J1/test-final/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestModle(t *testing.T){
	g:= NewGomegaWithT(t)

	t.Run(`student is required`,func(t *testing.T) {
		student := entity.Model{
			Student: "",
			FirstName:"gotji",
			Lastname:"gotji",
			Email:"siriphob@hotmail.com",
			Phone: "0981894780",
			Photo: "https://www.google.com",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(student)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(("studend is required")))
	})
}

func TestBestcase(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`All good`,func(t *testing.T){
		student := entity.Model{
			Student: "A12",
			FirstName:"gotji",
			Lastname:"gotji",
			Email:"siriphob@hotmail.com",
			Phone: "0981894780",
			Photo: "https://www.google.com",
			GenderID:  1,
		}

		ok,err := govalidator.ValidateStruct(student)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
		
	})
}