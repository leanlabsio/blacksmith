import {Component} from "angular2/core";
import {Inject} from "angular2/core";
import {Http} from "angular2/http";
import {RouteParams} from "angular2/router";
import {Headers} from "angular2/http";
import {ROUTER_DIRECTIVES} from "angular2/router";
import {Navigation} from "./../navigation/navigation";

export class Build {
    public username: string;
    public commit: string;
    public log:string;
    private _state:int;

    constructor(data) {
        this.username = data.username;
        this.commit = data.commit;
        this.log = data.log;
        this.state = data.state;
    }

    static create(data) {
        return new Build(data);
    }

    get state() {
        if (this._state == 2) {
            return "-";
        }
        if (this._state == 1) {
            return "+";
        }
        if (this._state == 0) {

        }
        return this._state;
    }

    set state(state:int) {
        this._state = state;
    }
}

@Component({
    selector: "build-list",
    template: <string>require('./build.list.html'),
    directives: [ROUTER_DIRECTIVES, Navigation],
})
export class BuildList
{
    public builds: Array<Build>;
    public repo: string;

    constructor(@Inject(Http) private http: Http, @Inject(RouteParams) private params: RouteParams){
        this.builds = [];
        this.repo = params.get("repo");
        var hs = new Headers();
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        http.get("/api/builds/"+this.repo, {headers:hs})
            .map(res => res.json())
            .subscribe(val => {
                if (val) {
                    val.forEach((b) => {
                        this.builds.push(Build.create(b));
                    })
                }
            });
    }
}
