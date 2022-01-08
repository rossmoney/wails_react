interface go {
  "main": {
    "App": {
		GetDepartments():Promise<string>
		Greet():Promise<string>
		Install(arg1:string):Promise<string>
		Sysvars():Promise<string>
    },
  }

  "sys": {
    "Sys": {
		All():Promise<Sys>
		Exec(arg1:string):Promise<void>
		GetRemoteFile(arg1:string,arg2:string):Promise<void>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
