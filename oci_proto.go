package oci

type OCIHandle uintptr
type OCISvcCtx uintptr
type OCIEnv uintptr
type OCIError uintptr

// oci api calls

//sys	OCIEnvCreate(envhpp *OCIEnv, mode uint32, ctxp uintptr, malocfp uintptr, ralocfp uintptr, mfreefp uintptr, xtramemsz uint32, usrmempp *uintptr) (err int16) = oci.OCIEnvCreate
//sys	OCIHandleFree(handle *OCIHandle, type uint32) (err int16) = oci.OCIHandleFree
//sys	OCILogon(envhp OCIEnv, errhp *OCIError, svchp *OCISvcCtx, username string, uname_len uint32, password string, passwd_len uint32, database string, dbname_len uint32) (err int16) = oci.OCILogon
