export namespace fuzzy {
	
	export class Rank {
	    Source: string;
	    Target: string;
	    Distance: number;
	    OriginalIndex: number;
	
	    static createFrom(source: any = {}) {
	        return new Rank(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Source = source["Source"];
	        this.Target = source["Target"];
	        this.Distance = source["Distance"];
	        this.OriginalIndex = source["OriginalIndex"];
	    }
	}

}

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
	export class ForManualReview {
	    submitted_name: string;
	    matches: fuzzy.Rank[];
	
	    static createFrom(source: any = {}) {
	        return new ForManualReview(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.submitted_name = source["submitted_name"];
	        this.matches = this.convertValues(source["matches"], fuzzy.Rank);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

