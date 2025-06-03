export namespace name_verification {
	
	export class CrimName {
	    fname: string;
	    mname: string;
	    lname: string;
	
	    static createFrom(source: any = {}) {
	        return new CrimName(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fname = source["fname"];
	        this.mname = source["mname"];
	        this.lname = source["lname"];
	    }
	}

}

