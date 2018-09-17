package sql_builder

import (
	"testing"
)

type Skill struct {
	Condition string `json:"condition" db:"condition"`
	Desc      string `json:"desc" db:"desc"`
	SkillType int    `json:"skillType" db:"skill_type"`
}

func TestSelectSqlBuilder_SelectByStruct(t *testing.T) {
	skill := Skill{
		Condition: "",
		SkillType: 1,
	}

	builder := Select("skill")
	builder.SelectByStruct(skill, true)

	if builder.String() != "SELECT `condition`,`desc`,`skill_type` FROM `skill` WHERE `skill_type`=?;" {
		t.Fatalf("Error -- sql string: %s \n", builder.String())
	}
}

func TestSelect(t *testing.T) {
	builder := Select("skill")
	builder.SetSearchFields("`condition`,`skill_type`")
	builder.WhereEq("skill_type", 1)
	builder.WhereIn("skill_type", []interface{}{1, 2})
	builder.WhereGt("skill_type", 1)
	builder.WhereLt("skill_type", 1)
	builder.WhereLike("condition", "vic")
	builder.WhereOr([]WhereOrCondition{
		{
			FieldName:  "skill_type",
			WhereType:  WHERE_TYPE_EQ,
			FieldValue: 1,
		},
	})
	if builder.String() != "SELECT `condition`,`skill_type` FROM `skill` WHERE (`skill_type`=? AND skill_type IN (?,?) AND `skill_type`>? AND `skill_type`<? AND `condition` LIKE ?) OR (`skill_type`=?);" {
		t.Fatalf("Error -- sql string: %s \n", builder.String())
	}
}

func TestUpdateSqlBuilder_UpdateByStruct(t *testing.T) {
	skill := Skill{
		Condition: "",
		SkillType: 1,
	}

	builder := Update("skill")
	builder.UpdateByStruct(skill, true)
	builder.WhereEq("skill_type", 1)
	builder.WhereIn("skill_type", []interface{}{1, 2})
	builder.WhereGt("skill_type", 1)
	builder.WhereLt("skill_type", 1)
	builder.WhereLike("condition", "vic")
	builder.WhereOr([]WhereOrCondition{
		{
			FieldName:  "skill_type",
			WhereType:  WHERE_TYPE_LT,
			FieldValue: 5,
		},
		{
			FieldName:  "skill_type",
			WhereType:  WHERE_TYPE_GT,
			FieldValue: 1,
		},
	})

	if builder.String() != "UPDATE `skill` SET `skill_type`=? WHERE (`skill_type`=? AND skill_type IN (?,?) AND `skill_type`>? AND `skill_type`<? AND `condition` LIKE ?) OR (`skill_type`<? AND `skill_type`>?);" {
		t.Fatalf("Error -- sql string: %s \n", builder.String())
	}
}

func TestInsertSqlBuilder_InsertByStruct(t *testing.T) {
	skill := Skill{
		SkillType: 1,
	}

	builder := Insert("skill")
	builder.InsertByStruct(skill)

	if builder.String() != "INSERT INTO `skill`(condition,desc,skill_type) VALUES(?,?,?);" {
		t.Fatalf("Error -- sql string: %s \n", builder.String())
	}
}

func TestDeleteSqlBuilder_DeleteAndWhereByStruct(t *testing.T) {
	skill := Skill{
		SkillType: 1,
	}

	builder := Delete("skill")
	builder.WhereByStruct(skill, true)

	if builder.String() != "DELETE FROM `skill` WHERE `skill_type`=?;" {
		t.Fatalf("Error -- sql string: %s \n", builder.String())
	}
}
