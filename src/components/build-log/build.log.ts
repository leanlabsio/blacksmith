import {
    Component,
    Inject,
    OnInit
} from "angular2/core";

import {
    Http,
    Headers
} from "angular2/http";

import {RouteParams} from "angular2/router";
import {Build} from "./../build-list/build.list.ts";
import {NAVIGATION_DIRECTIVES} from "./../navigation/navigation";

const template: string = <string>require('./build.log.html');

@Component({
    template: template,
    directives: [NAVIGATION_DIRECTIVES]
})
export class BuildLog implements OnInit
{
    public build:Build;

    constructor(@Inject(Http) private http: Http, @Inject(RouteParams) private params: RouteParams) {
        let repo = params.get("repo");
        let commit = params.get("commit");
        let hs = new Headers();
        hs.append("Authorization", "Bearer "+localStorage.getItem("jwt"));
        this.http.get("/api/logs/"+repo+"?commit="+commit, {headers:hs})
            .map(res => res.json())
            .subscribe(val => this.build = Build.create(val));
    }

    ngOnInit() {
        this.build = new Build({});
    }
}