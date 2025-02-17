export namespace main {
	
	export class ToDo {
	    Task: string;
	    Done: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ToDo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Task = source["Task"];
	        this.Done = source["Done"];
	    }
	}

}

