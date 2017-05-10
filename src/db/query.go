/**
 * Created by 93201 on 2017/5/10.
 */
package db

type qStruct struct {
	Page  int    `form:"page"`
	Count int    `form:"count"`
	Order string `form:"-"`
}
