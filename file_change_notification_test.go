// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-23 11:40:22 1B149C      [zr-win/file_change_notification_test.go]
// -----------------------------------------------------------------------------

package win

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/zr" // Zircon-Go

/*
to test all items in file_change_notification.go use:
    go test --run Test_fchn_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

// go test --run Test_fchn_FindFirstChangeNotification_
func Test_fchn_FindFirstChangeNotification_(t *testing.T) {
	zr.TBegin(t)
	// FindFirstChangeNotification(
	//      lpPathName string,
	//      bWatchSubtree BOOL,
	//      dwNotifyFilter DWORD,
	//  ) HANDLE
	//
	// create a subfolder called TEST
	// write a timestamp to file A
	// wait for one second
	// write a timestamp to file B
	//
	var NOTIFY = FILE_NOTIFY_CHANGE_FILE_NAME |
		FILE_NOTIFY_CHANGE_DIR_NAME |
		FILE_NOTIFY_CHANGE_ATTRIBUTES |
		FILE_NOTIFY_CHANGE_SIZE |
		FILE_NOTIFY_CHANGE_LAST_WRITE |
		FILE_NOTIFY_CHANGE_LAST_ACCESS |
		FILE_NOTIFY_CHANGE_CREATION |
		FILE_NOTIFY_CHANGE_SECURITY
	//
	// call with invalid folder:
	// should return INVALID_HANDLE_VALUE i.e. -1 / 0xFFFFFFFFFFFFFFFF
	var got = HANDLE(FindFirstChangeNotification(`z:\yx`, TRUE, DWORD(NOTIFY)))
	zr.TTrue(t, (got&INVALID_HANDLE_VALUE) == INVALID_HANDLE_VALUE)
	//
	// call with valid folder:
	got = FindFirstChangeNotification(".", TRUE, DWORD(NOTIFY))
	zr.TTrue(t, got != 0)
	zr.TTrue(t, got < 65535)
} //                                      Test_fchn_FindFirstChangeNotification_

// go test --run Test_fchn_FileChangeNotifications_
func Test_fchn_FileChangeNotifications_(t *testing.T) {
	if true {
		PL("Test_fchn_FileChangeNotifications_ IS UNIFINISHED")
		return
	}
	const WAIT_OBJECT = 0 //TODO: move constants to github.com/balacode/zr-win
	const WAIT_OBJECT_0 = 0
	const WAIT_TIMEOUT = 258
	const NULL = 0
	const INFINITE = 0xFFFFFFFF // infinite timeout
	var path = `X:\TEST`
	var NOTIFY = FILE_NOTIFY_CHANGE_CREATION |
		FILE_NOTIFY_CHANGE_FILE_NAME |
		FILE_NOTIFY_CHANGE_LAST_WRITE |
		FILE_NOTIFY_CHANGE_SIZE |
		0
	// not used:
	// FILE_NOTIFY_CHANGE_ATTRIBUTES |
	// FILE_NOTIFY_CHANGE_DIR_NAME |
	// FILE_NOTIFY_CHANGE_LAST_ACCESS |
	// FILE_NOTIFY_CHANGE_SECURITY |
	var tms = zr.Timestamp
	//
	// start watching the folder
	var handles = []HANDLE{
		FindFirstChangeNotification(path, TRUE, DWORD(NOTIFY)),
	}
	// check that handle value is correct
	switch handles[0] {
	case INVALID_HANDLE_VALUE:
		fmt.Println(tms(), "FindFirstChangeNotification() failed")
		t.Fail()
		return
	case NULL:
		fmt.Println(tms(), "FindFirstChangeNotification() returned NULL")
		t.Fail()
		return
	default:
		fmt.Println(tms(), "FindFirstChangeNotification() handle:",
			handles[0])
	}
	for {
		// wait for notification
		fmt.Println(tms(), "Waiting.."+".")
		var status = WaitForMultipleObjects(1, &handles[0], TRUE, INFINITE)
		fmt.Println(tms(), "Wait ended"+".")
		switch status {
		case WAIT_OBJECT_0:
			fmt.Println(tms(), "DIRECTORY CHANGED!")
			if FindNextChangeNotification(handles[0]) == FALSE {
				fmt.Println(tms(), "FindNextChangeNotification() failed")
				t.Fail()
				return
			}
		case WAIT_OBJECT_0 + 1:
			// notifications from other handle(s) which we don't have here
		case WAIT_TIMEOUT:
			// A timeout occurred, this would happen if some value other
			// than INFINITE is used in the Wait call and no changes occur.
			// In a single-threaded environment you might not want an
			// INFINITE wait.
			fmt.Println(tms(), "No changes during timeout period")
		default:
			fmt.Println(tms(), "Unhandled wait status", status)
			t.Fail()
			return
		}
	}
} //                                          Test_fchn_FileChangeNotifications_

//end
