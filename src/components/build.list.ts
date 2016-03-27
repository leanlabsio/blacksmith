import {Component} from "angular2/core";
import {View} from "angular2/core";
import {Inject} from "angular2/core";
import {Http} from "angular2/http";
import {RouteParams} from "angular2/router";
import {Headers} from "angular2/http";
import {ROUTER_DIRECTIVES} from "angular2/router";

export class Build {
    public username: string;
    public commit: string;
    public log:string;

    constructor(data) {
        this.username = data.username;
        this.commit = data.commit;
        this.log = data.log;
    }

    static create(data) {
        return new Build(data);
    }
}

@Component({
    selector: "build-list",
    template: `
    <div class="row align-center" *ngIf="!builds || builds.length == 0">
        No builds yet
    </div>
    <div class="row align-center" *ngFor="#build of builds">
        <div class="columns medium-1"></div>
        <div class="columns medium-9">
            <a [routerLink]="['BuildLog', {repo: repo, commit: build.commit}]">
            {{build.commit}}
            </a>
            <br/>
            {{build.username}}
        </div>
        <div class="columns medium-2">
        status
        </div>
    </div>
    `,
    directives: [ROUTER_DIRECTIVES],
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
