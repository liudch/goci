package oci

import (
	"syscall"
	"unsafe"
)

const (
	OCI_DEFAULT  = 0
	OCI_THREADED = 1
	OCI_OBJECT   = 2
	OCI_EVENTS   = 4
	OCI_SHARED   = 16

	OCI_HTYPE_FIRST                = 1                 /* start value of handle type */
	OCI_HTYPE_ENV                  = 1                 /* environment handle */
	OCI_HTYPE_ERROR                = 2                 /* error handle */
	OCI_HTYPE_SVCCTX               = 3                 /* service handle */
	OCI_HTYPE_STMT                 = 4                 /* statement handle */
	OCI_HTYPE_BIND                 = 5                 /* bind handle */
	OCI_HTYPE_DEFINE               = 6                 /* define handle */
	OCI_HTYPE_DESCRIBE             = 7                 /* describe handle */
	OCI_HTYPE_SERVER               = 8                 /* server handle */
	OCI_HTYPE_SESSION              = 9                 /* authentication handle */
	OCI_HTYPE_AUTHINFO             = OCI_HTYPE_SESSION /* SessionGet auth handle */
	OCI_HTYPE_TRANS                = 10                /* transaction handle */
	OCI_HTYPE_COMPLEXOBJECT        = 11                /* complex object retrieval handle */
	OCI_HTYPE_SECURITY             = 12                /* security handle */
	OCI_HTYPE_SUBSCRIPTION         = 13                /* subscription handle */
	OCI_HTYPE_DIRPATH_CTX          = 14                /* direct path context */
	OCI_HTYPE_DIRPATH_COLUMN_ARRAY = 15                /* direct path column array */
	OCI_HTYPE_DIRPATH_STREAM       = 16                /* direct path stream */
	OCI_HTYPE_PROC                 = 17                /* process handle */
	OCI_HTYPE_DIRPATH_FN_CTX       = 18                /* direct path function context */
	OCI_HTYPE_DIRPATH_FN_COL_ARRAY = 19                /* dp object column array */
	OCI_HTYPE_XADSESSION           = 20                /* access driver session */
	OCI_HTYPE_XADTABLE             = 21                /* access driver table */
	OCI_HTYPE_XADFIELD             = 22                /* access driver field */
	OCI_HTYPE_XADGRANULE           = 23                /* access driver granule */
	OCI_HTYPE_XADRECORD            = 24                /* access driver record */
	OCI_HTYPE_XADIO                = 25                /* access driver I/O */
	OCI_HTYPE_CPOOL                = 26                /* connection pool handle */
	OCI_HTYPE_SPOOL                = 27                /* session pool handle */
	OCI_HTYPE_ADMIN                = 28                /* admin handle */
	OCI_HTYPE_EVENT                = 29                /* HA event handle */
	OCI_HTYPE_LAST                 = 29                /* last value of a handle type */

)

var (
	modoci = syscall.NewLazyDLL("C:\\Program Files\\Oracle\\instantclient_11_2\\oci.dll")

	// Connect, authorize, and initialize
	procOCIAppCtxClearAll        = modoci.NewProc("OCIAppCtxClearAll")        // Clear all attribute-value information in a namespace of an application context
	procOCIAppCtxSet             = modoci.NewProc("OCIAppCtxSet")             // Set an attribute and its associated value in a namespace of an application context
	procOCIConnectionPoolCreate  = modoci.NewProc("OCIConnectionPoolCreate")  // Initialize the connection pool
	procOCIConnectionPoolDestroy = modoci.NewProc("OCIConnectionPoolDestroy") // Destroy the connection pool
	procOCIDBShutdown            = modoci.NewProc("OCIDBShutdown")            // Shut down Oracle Database
	procOCIDBStartup             = modoci.NewProc("OCIDBStartup")             // Start an Oracle Database instance
	procOCIEnvCreate             = modoci.NewProc("OCIEnvCreate")             // Create and initialize an OCI environment handle
	procOCIEnvNlsCreate          = modoci.NewProc("OCIEnvNlsCreate")          // Create and initialize an environment handle for OCI functions to work under. Enable you to set character set ID and national character set ID at environment creation time.
	procOCILogoff                = modoci.NewProc("OCILogoff")                // Release a session that was retrieved using OCILogon2() or OCILogon()
	procOCILogon                 = modoci.NewProc("OCILogon")                 // Simplify single-session logon
	procOCILogon2                = modoci.NewProc("OCILogon2")                // Create a logon session in various modes
	procOCIServerAttach          = modoci.NewProc("OCIServerAttach")          // Attach to a server; initialize server context handle
	procOCIServerDetach          = modoci.NewProc("OCIServerDetach")          // Detach from a server; uninitialize server context handle
	procOCISessionBegin          = modoci.NewProc("OCISessionBegin")          // Authenticate a user
	procOCISessionEnd            = modoci.NewProc("OCISessionEnd")            // Terminate a user session
	procOCISessionGet            = modoci.NewProc("OCISessionGet")            // Get a session from a session pool
	procOCISessionPoolCreate     = modoci.NewProc("OCISessionPoolCreate")     // Initialize a session pool
	procOCISessionPoolDestroy    = modoci.NewProc("OCISessionPoolDestroy")    // Destroy a session pool
	procOCISessionRelease        = modoci.NewProc("OCISessionRelease")        // Release a session
	procOCITerminate             = modoci.NewProc("OCITerminate")             // Detach from a shared memory subsystem

	// Statement
	procOCIStmtExecute      = modoci.NewProc("OCIStmtExecute")      // Send statements to server for execution
	procOCIStmtFetch2       = modoci.NewProc("OCIStmtFetch2")       // Fetch rows from a query and fetches a row from the (scrollable) result set.
	procOCIStmtGetPieceInfo = modoci.NewProc("OCIStmtGetPieceInfo") // Get piece information for piecewise operations
	procOCIStmtPrepare      = modoci.NewProc("OCIStmtPrepare")      // Prepare a SQL or PL/SQL statement for execution
	procOCIStmtPrepare2     = modoci.NewProc("OCIStmtPrepare2")     // Prepare a SQL or PL/SQL statement for execution. The user also has the option of using the statement cache, if it has been enabled.
	procOCIStmtRelease      = modoci.NewProc("OCIStmtRelease")      // Release the statement handle
	procOCIStmtSetPieceInfo = modoci.NewProc("OCIStmtSetPieceInfo") // Set piece information for piecewise operations

	// Bind, define, and describe
	procOCIBindArrayOfStruct   = modoci.NewProc("OCIBindArrayOfStruct")   // Set skip parameters for static array bind
	procOCIBindByName          = modoci.NewProc("OCIBindByName")          // Bind by name
	procOCIBindByPos           = modoci.NewProc("OCIBindByPos")           // Bind by position
	procOCIBindDynamic         = modoci.NewProc("OCIBindDynamic")         // Set additional attributes after bind with OCI_DATA_AT_EXEC mode
	procOCIBindObject          = modoci.NewProc("OCIBindObject")          // Set additional attributes for bind of named data type
	procOCIDefineArrayOfStruct = modoci.NewProc("OCIDefineArrayOfStruct") // Set additional attributes for static array define
	procOCIDefineByPos         = modoci.NewProc("OCIDefineByPos")         // Define an output variable association
	procOCIDefineDynamic       = modoci.NewProc("OCIDefineDynamic")       // Set additional attributes for define in OCI_DYNAMIC_FETCH mode
	procOCIDefineObject        = modoci.NewProc("OCIDefineObject")        // Set additional attributes for define of named data type
	procOCIDescribeAny         = modoci.NewProc("OCIDescribeAny")         // Describe existing schema objects
	procOCIStmtGetBindInfo     = modoci.NewProc("OCIStmtGetBindInfo")     // Get bind and indicator variable names and handle

	// Handle and descriptor
	procOCIArrayDescriptorAlloc = modoci.NewProc("OCIArrayDescriptorAlloc") // Allocate an array of descriptors
	procOCIArrayDescriptorFree  = modoci.NewProc("OCIArrayDescriptorFree")  // Free an array of descriptors
	procOCIAttrGet              = modoci.NewProc("OCIAttrGet")              // Get the value of an attribute of a handle
	procOCIAttrSet              = modoci.NewProc("OCIAttrSet")              // Set the value of an attribute of a handle or descriptor
	procOCIDescriptorAlloc      = modoci.NewProc("OCIDescriptorAlloc")      // Allocate and initialize a descriptor or LOB locator
	procOCIDescriptorFree       = modoci.NewProc("OCIDescriptorFree")       // Free a previously allocated descriptor
	procOCIHandleAlloc          = modoci.NewProc("OCIHandleAlloc")          // Allocate and initialize a handle
	procOCIHandleFree           = modoci.NewProc("OCIHandleFree")           // Free a previously allocated handle
	procOCIParamGet             = modoci.NewProc("OCIParamGet")             // Get a parameter descriptor
	procOCIParamSet             = modoci.NewProc("OCIParamSet")             // Set parameter descriptor in COR handle

	// Direct path loading
	procOCIDirPathAbort            = modoci.NewProc("OCIDirPathAbort")            // Terminate a direct path operation
	procOCIDirPathColArrayEntryGet = modoci.NewProc("OCIDirPathColArrayEntryGet") // Get a specified entry in a column array
	procOCIDirPathColArrayEntrySet = modoci.NewProc("OCIDirPathColArrayEntrySet") // Set a specified entry in a column array to a specific value
	procOCIDirPathColArrayReset    = modoci.NewProc("OCIDirPathColArrayReset")    // Reset the row array state
	procOCIDirPathColArrayRowGet   = modoci.NewProc("OCIDirPathColArrayRowGet")   // Get the base row pointers for a specified row number
	procOCIDirPathColArrayToStream = modoci.NewProc("OCIDirPathColArrayToStream") // Convert from a column array to a direct path stream format
	procOCIDirPathDataSave         = modoci.NewProc("OCIDirPathDataSave")         // Do a data savepoint, or commit the loaded data and finish the load operation
	procOCIDirPathFinish           = modoci.NewProc("OCIDirPathFinish")           // Finish and commit the loaded data
	procOCIDirPathFlushRow         = modoci.NewProc("OCIDirPathFlushRow")         // Deprecated.
	procOCIDirPathLoadStream       = modoci.NewProc("OCIDirPathLoadStream")       // Load the data converted to direct path stream format
	procOCIDirPathPrepare          = modoci.NewProc("OCIDirPathPrepare")          // Prepare direct path interface to convert or load rows
	procOCIDirPathStreamReset      = modoci.NewProc("OCIDirPathStreamReset")      // Reset the direct path stream state

	// Collection and iterator
	procOCICollAppend       = modoci.NewProc("OCICollAppend")       // Append an element to the end of a collection
	procOCICollAssign       = modoci.NewProc("OCICollAssign")       // Assign (deep copy) one collection to another
	procOCICollAssignElem   = modoci.NewProc("OCICollAssignElem")   // Assign the given element value elem to the element at coll[index]
	procOCICollGetElem      = modoci.NewProc("OCICollGetElem")      // Get pointer to an element
	procOCICollGetElemArray = modoci.NewProc("OCICollGetElemArray") // Get an array of elements from a collection
	procOCICollIsLocator    = modoci.NewProc("OCICollIsLocator")    // Indicate whether a collection is locator-based or not
	procOCICollMax          = modoci.NewProc("OCICollMax")          // Return maximum number of elements in collection
	procOCICollSize         = modoci.NewProc("OCICollSize")         // Get current size of collection (in number of elements)
	procOCICollTrim         = modoci.NewProc("OCICollTrim")         // Trim elements from the collection
	procOCIIterCreate       = modoci.NewProc("OCIIterCreate")       // Create iterator to scan the varray elements
	procOCIIterDelete       = modoci.NewProc("OCIIterDelete")       // Delete iterator
	procOCIIterGetCurrent   = modoci.NewProc("OCIIterGetCurrent")   // Get current collection element
	procOCIIterInit         = modoci.NewProc("OCIIterInit")         // Initialize iterator to scan the given collection
	procOCIIterNext         = modoci.NewProc("OCIIterNext")         // Get next collection element
	procOCIIterPrev         = modoci.NewProc("OCIIterPrev")         // Get previous collection element	

	// Date, date time, and interval
	procOCIDateAddDays               = modoci.NewProc("OCIDateAddDays")               // Add or subtract days
	procOCIDateAddMonths             = modoci.NewProc("OCIDateAddMonths")             // Add or subtract months
	procOCIDateAssign                = modoci.NewProc("OCIDateAssign")                // Assign date
	procOCIDateCheck                 = modoci.NewProc("OCIDateCheck")                 // Check if the given date is valid
	procOCIDateCompare               = modoci.NewProc("OCIDateCompare")               // Compare dates
	procOCIDateDaysBetween           = modoci.NewProc("OCIDateDaysBetween")           // Get number of days between two dates
	procOCIDateFromText              = modoci.NewProc("OCIDateFromText")              // Convert string to date
	procOCIDateGetDate               = modoci.NewProc("OCIDateGetDate")               // Get the date portion of a date
	procOCIDateGetTime               = modoci.NewProc("OCIDateGetTime")               // Get the time portion of a date
	procOCIDateLastDay               = modoci.NewProc("OCIDateLastDay")               // Get date of last day of month
	procOCIDateNextDay               = modoci.NewProc("OCIDateNextDay")               // Get date of next day
	procOCIDateSetDate               = modoci.NewProc("OCIDateSetDate")               // Set the date portion of a date
	procOCIDateSetTime               = modoci.NewProc("OCIDateSetTime")               // Set the time portion of a date
	procOCIDateSysDate               = modoci.NewProc("OCIDateSysDate")               // Get the current system date and time
	procOCIDateTimeAssign            = modoci.NewProc("OCIDateTimeAssign")            // Perform a datetime assignment
	procOCIDateTimeCheck             = modoci.NewProc("OCIDateTimeCheck")             // Check if the given date is valid
	procOCIDateTimeCompare           = modoci.NewProc("OCIDateTimeCompare")           // Compare two datetime values
	procOCIDateTimeConstruct         = modoci.NewProc("OCIDateTimeConstruct")         // Construct a datetime descriptor
	procOCIDateTimeConvert           = modoci.NewProc("OCIDateTimeConvert")           // Convert one datetime type to another
	procOCIDateTimeFromArray         = modoci.NewProc("OCIDateTimeFromArray")         // Convert an array of size OCI_DT_ARRAYLEN to an OCIDateTime descriptor
	procOCIDateTimeFromText          = modoci.NewProc("OCIDateTimeFromText")          // Convert the given string to Oracle datetime type in the OCIDateTime descriptor, according to the specified format
	procOCIDateTimeGetDate           = modoci.NewProc("OCIDateTimeGetDate")           // Get the date (year, month, day) portion of a datetime value
	procOCIDateTimeGetTime           = modoci.NewProc("OCIDateTimeGetTime")           // Get the time (hour, min, second, fractional second) of a datetime value
	procOCIDateTimeGetTimeZoneName   = modoci.NewProc("OCIDateTimeGetTimeZoneName")   // Get the time zone name portion of a datetime value
	procOCIDateTimeGetTimeZoneOffset = modoci.NewProc("OCIDateTimeGetTimeZoneOffset") // Get the time zone (hour, minute) portion of a datetime value
	procOCIDateTimeIntervalAdd       = modoci.NewProc("OCIDateTimeIntervalAdd")       // Add an interval to a datetime to produce a resulting datetime
	procOCIDateTimeIntervalSub       = modoci.NewProc("OCIDateTimeIntervalSub")       // Subtract an interval from a datetime and store the result in a datetime
	procOCIDateTimeSubtract          = modoci.NewProc("OCIDateTimeSubtract")          // Take two datetimes as input and store their difference in an interval
	procOCIDateTimeSysTimeStamp      = modoci.NewProc("OCIDateTimeSysTimeStamp")      // Get the system current date and time as a time stamp with time zone
	procOCIDateTimeToArray           = modoci.NewProc("OCIDateTimeToArray")           // Convert an OCIDateTime descriptor to an array
	procOCIDateTimeToText            = modoci.NewProc("OCIDateTimeToText")            // Convert the given date to a string according to the specified format
	procOCIDateToText                = modoci.NewProc("OCIDateToText")                // Convert date to string
	procOCIDateZoneToZone            = modoci.NewProc("OCIDateZoneToZone")            // Convert date from one time zone to another zone
	procOCIIntervalAdd               = modoci.NewProc("OCIIntervalAdd")               // Add two intervals to produce a resulting interval
	procOCIIntervalAssign            = modoci.NewProc("OCIIntervalAssign")            // Copy one interval to another
	procOCIIntervalCheck             = modoci.NewProc("OCIIntervalCheck")             // Check the validity of an interval
	procOCIIntervalCompare           = modoci.NewProc("OCIIntervalCompare")           // Compare two intervals
	procOCIIntervalDivide            = modoci.NewProc("OCIIntervalDivide")            // Divide an interval by an Oracle NUMBER to produce an interval
	procOCIIntervalFromNumber        = modoci.NewProc("OCIIntervalFromNumber")        // Convert an Oracle NUMBER to an interval
	procOCIIntervalFromText          = modoci.NewProc("OCIIntervalFromText")          // When given an interval string, return the interval represented by the string
	procOCIIntervalFromTZ            = modoci.NewProc("OCIIntervalFromTZ")            // Return an OCI_DTYPE_INTERVAL_DS
	procOCIIntervalGetDaySecond      = modoci.NewProc("OCIIntervalGetDaySecond")      // Get values of day, hour, minute, and second from an interval
	procOCIIntervalGetYearMonth      = modoci.NewProc("OCIIntervalGetYearMonth")      // Get year and month from an interval
	procOCIIntervalMultiply          = modoci.NewProc("OCIIntervalMultiply")          // Multiply an interval by an Oracle NUMBER to produce an interval
	procOCIIntervalSetDaySecond      = modoci.NewProc("OCIIntervalSetDaySecond")      // Set day, hour, minute, and second in an interval
	procOCIIntervalSetYearMonth      = modoci.NewProc("OCIIntervalSetYearMonth")      // Set year and month in an interval
	procOCIIntervalSubtract          = modoci.NewProc("OCIIntervalSubtract")          // Subtract two intervals and stores the result in an interval
	procOCIIntervalToNumber          = modoci.NewProc("OCIIntervalToNumber")          // Convert an interval to an Oracle NUMBER
	procOCIIntervalToText            = modoci.NewProc("OCIIntervalToText")            // When given an interval, produce a string representing the interval

	// Numeric
	procOCINumberAbs         = modoci.NewProc("OCINumberAbs")         // Compute the absolute value
	procOCINumberAdd         = modoci.NewProc("OCINumberAdd")         // Add NUMBERs
	procOCINumberArcCos      = modoci.NewProc("OCINumberArcCos")      // Compute the arc cosine
	procOCINumberArcSin      = modoci.NewProc("OCINumberArcSin")      // Compute the arc sine
	procOCINumberArcTan      = modoci.NewProc("OCINumberArcTan")      // Compute the arc tangent
	procOCINumberArcTan2     = modoci.NewProc("OCINumberArcTan2")     // Compute the arc tangent of two NUMBERs
	procOCINumberAssign      = modoci.NewProc("OCINumberAssign")      // Assign one NUMBER to another
	procOCINumberCeil        = modoci.NewProc("OCINumberCeil")        // Compute the ceiling of NUMBER
	procOCINumberCmp         = modoci.NewProc("OCINumberCmp")         // Compare NUMBERs
	procOCINumberCos         = modoci.NewProc("OCINumberCos")         // Compute the cosine
	procOCINumberDec         = modoci.NewProc("OCINumberDec")         // Decrement a NUMBER
	procOCINumberDiv         = modoci.NewProc("OCINumberDiv")         // Divide two NUMBERs
	procOCINumberExp         = modoci.NewProc("OCINumberExp")         // Raise e to the specified Oracle NUMBER power
	procOCINumberFloor       = modoci.NewProc("OCINumberFloor")       // Compute the floor value of a NUMBER
	procOCINumberFromInt     = modoci.NewProc("OCINumberFromInt")     // Convert an integer to an Oracle NUMBER
	procOCINumberFromReal    = modoci.NewProc("OCINumberFromReal")    // Convert a real type to an Oracle NUMBER
	procOCINumberFromText    = modoci.NewProc("OCINumberFromText")    // Convert a string to an Oracle NUMBER
	procOCINumberHypCos      = modoci.NewProc("OCINumberHypCos")      // Compute the hyperbolic cosine
	procOCINumberHypSin      = modoci.NewProc("OCINumberHypSin")      // Compute the hyperbolic sine
	procOCINumberHypTan      = modoci.NewProc("OCINumberHypTan")      // Compute the hyperbolic tangent
	procOCINumberInc         = modoci.NewProc("OCINumberInc")         // Increment an Oracle NUMBER
	procOCINumberIntPower    = modoci.NewProc("OCINumberIntPower")    // Raise a given base to an integer power
	procOCINumberIsInt       = modoci.NewProc("OCINumberIsInt")       // Test if a NUMBER is an integer
	procOCINumberIsZero      = modoci.NewProc("OCINumberIsZero")      // Test if a NUMBER is zero
	procOCINumberLn          = modoci.NewProc("OCINumberLn")          // Compute the natural logarithm
	procOCINumberLog         = modoci.NewProc("OCINumberLog")         // Compute the logarithm to an arbitrary base
	procOCINumberMod         = modoci.NewProc("OCINumberMod")         // Gets the modulus (remainder) of the division of two Oracle NUMBERs
	procOCINumberMul         = modoci.NewProc("OCINumberMul")         // Multiply two Oracle NUMBERs
	procOCINumberNeg         = modoci.NewProc("OCINumberNeg")         // Negates an Oracle NUMBER
	procOCINumberPower       = modoci.NewProc("OCINumberPower")       // Raises a given base to a given exponent
	procOCINumberPrec        = modoci.NewProc("OCINumberPrec")        // Round a NUMBER to a specified number of decimal places
	procOCINumberRound       = modoci.NewProc("OCINumberRound")       // Round an Oracle NUMBER to a specified decimal place
	procOCINumberSetPi       = modoci.NewProc("OCINumberSetPi")       // Initialize a NUMBER to pi
	procOCINumberSetZero     = modoci.NewProc("OCINumberSetZero")     // Initialize a NUMBER to zero
	procOCINumberShift       = modoci.NewProc("OCINumberShift")       // Multiply by 10, shifting a specified number of decimal places
	procOCINumberSign        = modoci.NewProc("OCINumberSign")        // Obtain the sign of an Oracle NUMBER
	procOCINumberSin         = modoci.NewProc("OCINumberSin")         // Compute the sine
	procOCINumberSqrt        = modoci.NewProc("OCINumberSqrt")        // Compute the square root of a NUMBER
	procOCINumberSub         = modoci.NewProc("OCINumberSub")         // Subtract NUMBERs
	procOCINumberTan         = modoci.NewProc("OCINumberTan")         // Compute the tangent
	procOCINumberToInt       = modoci.NewProc("OCINumberToInt")       // Convert an Oracle NUMBER to an integer
	procOCINumberToReal      = modoci.NewProc("OCINumberToReal")      // Convert an Oracle NUMBER to a real type
	procOCINumberToRealArray = modoci.NewProc("OCINumberToRealArray") // Convert an array of NUMBER to a real array.
	procOCINumberToText      = modoci.NewProc("OCINumberToText")      // Convert an Oracle NUMBER to a string
	procOCINumberTrunc       = modoci.NewProc("OCINumberTrunc")       // Truncate an Oracle NUMBER at a specified decimal place	

	// Raw
	procOCIRawAllocSize   = modoci.NewProc("OCIRawAllocSize")   // Get allocated size of raw memory in bytes
	procOCIRawAssignBytes = modoci.NewProc("OCIRawAssignBytes") // Assign raw bytes to raw
	procOCIRawAssignRaw   = modoci.NewProc("OCIRawAssignRaw")   // Assign raw to raw
	procOCIRawPtr         = modoci.NewProc("OCIRawPtr")         // Get raw data pointer
	procOCIRawResize      = modoci.NewProc("OCIRawResize")      // Resize memory of variable-length raw
	procOCIRawSize        = modoci.NewProc("OCIRawSize")        // Get raw size	

	// Ref
	procOCIRefAssign  = modoci.NewProc("OCIRefAssign")  // Assign one REF to another
	procOCIRefClear   = modoci.NewProc("OCIRefClear")   // Clear or nullify a REF
	procOCIRefFromHex = modoci.NewProc("OCIRefFromHex") // Convert hexadecimal string to REF
	procOCIRefHexSize = modoci.NewProc("OCIRefHexSize") // Return size of hexadecimal representation of REF
	procOCIRefIsEqual = modoci.NewProc("OCIRefIsEqual") // Compare two REFs for equality
	procOCIRefIsNull  = modoci.NewProc("OCIRefIsNull")  // Test if a REF is NULL
	procOCIRefToHex   = modoci.NewProc("OCIRefToHex")   // Convert REF to hexadecimal string

	// String
	procOCIStringAllocSize  = modoci.NewProc("OCIStringAllocSize")  // Get the allocated size of string memory in bytes
	procOCIStringAssign     = modoci.NewProc("OCIStringAssign")     // Assign a string to a string
	procOCIStringAssignText = modoci.NewProc("OCIStringAssignText") // Assign a text string to a string
	procOCIStringPtr        = modoci.NewProc("OCIStringPtr")        // Get a string pointer
	procOCIStringResize     = modoci.NewProc("OCIStringResize")     // Resize the string memory
	procOCIStringSize       = modoci.NewProc("OCIStringSize")       // Get the string size

	// Table
	procOCITableDelete = modoci.NewProc("OCITableDelete") // Delete element
	procOCITableExists = modoci.NewProc("OCITableExists") // Test whether element exists
	procOCITableFirst  = modoci.NewProc("OCITableFirst")  // Return first index of table
	procOCITableLast   = modoci.NewProc("OCITableLast")   // Return last index of table
	procOCITableNext   = modoci.NewProc("OCITableNext")   // Return next available index of table
	procOCITablePrev   = modoci.NewProc("OCITablePrev")   // Return previous available index of table
	procOCITableSize   = modoci.NewProc("OCITableSize")   // Return current size of table
)

type OCIHandle struct {
	h uintptr // Handle
	t uint    // Handle type
}

func OCIEnvCreate(mode uint32) (henv *OCIHandle, err error) {
	var e *OCIHandle = new(OCIHandle) // Environment handle
	e.t = OCI_HTYPE_ENV

	r0, _, e1 := procOCIEnvCreate.Call( // sword OCIEnvCreate   (
		uintptr(unsafe.Pointer(&e.h)), // OCIEnv        **envhpp,
		uintptr(mode),                 // ub4           mode,
		0,                             // CONST dvoid   *ctxp,
		0,                             // CONST dvoid   *(*malocfp) (dvoid *ctxp, size_t size),
		0,                             // CONST dvoid   *(*ralocfp) (dvoid *ctxp, dvoid *memptr, size_t newsize),
		0,                             // CONST void    (*mfreefp) (dvoid *ctxp, dvoid *memptr))
		0,                             // size_t        xtramemsz,
		0)                             // dvoid         **usrmempp );

	if r0 != 0 {
		return nil, error(e1)
	}
	return e, nil
}

func OCIHandleFree(handle *OCIHandle) (err error) {
	r0, _, e1 := procOCIHandleFree.Call( // sword OCIHandleFree (
		handle.h,          // dvoid     *hndlp,
		uintptr(handle.t)) // ub4       type );
	if r0 != 0 {
		err = error(e1)
	} else {
		err = nil
	}
	return
}

func OCILogon(env *OCIHandle, username, password, database string) (svcctx *OCIHandle, err error) {
	var s *OCIHandle = new(OCIHandle) // Service context
	s.t = OCI_HTYPE_SVCCTX
	var e *OCIHandle = new(OCIHandle) // Error handle
	e.t = OCI_HTYPE_ERROR

	r0, _, e1 := procOCILogon.Call( // sword OCILogon (
		env.h, // OCIEnv          *envhp,
		uintptr(unsafe.Pointer(&e.h)),                               // OCIError        *errhp,
		uintptr(unsafe.Pointer(&s.h)),                               // OCISvcCtx       **svchp,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(username))), // CONST OraText   *username,
		uintptr(len(username)),                                      // ub4             uname_len,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(password))), // CONST OraText   *password,
		uintptr(len(password)),                                      // ub4             passwd_len,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(database))), // CONST OraText   *dbname,
		uintptr(len(database)))                                      // ub4             dbname_len );

	if r0 != 0 {
		err = error(e1)
		return
	}

	return s, nil
}
