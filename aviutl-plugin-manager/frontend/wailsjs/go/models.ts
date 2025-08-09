export namespace library {

	export class Plugin {
	    id: string;
	    name: string;
	    version: string;
	    author: string;
	    description: string;
	    files: string[];

	    static createFrom(source: any = {}) {
	        return new Plugin(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.version = source["version"];
	        this.author = source["author"];
	        this.description = source["description"];
	        this.files = source["files"];
	    }
	}

}

export namespace profile {

	export class Profile {
	    name: string;
	    enabled_plugins: Record<string, boolean>;

	    static createFrom(source: any = {}) {
	        return new Profile(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.enabled_plugins = source["enabled_plugins"];
	    }
	}

}
