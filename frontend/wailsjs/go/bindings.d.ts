interface go {
  "main": {
    "App": {
		Greet():Promise<string>
		Sysvars():Promise<string>
    },
  }

  "sys": {
    "Sys": {
		All():Promise<Sys>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
