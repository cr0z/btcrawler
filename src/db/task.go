/**
 * Created by 93201 on 2017/5/10.
 */
package db

type task interface {
	init()
	start()
	onCall(interface{})
	stop()
}
