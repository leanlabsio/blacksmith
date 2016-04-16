import {Component} from "angular2/core";
import {Inject} from "angular2/core";
import {RouteParams} from "angular2/router";
import {Http} from "angular2/http";
import {Headers} from "angular2/http";
import {Build} from "./../build-list/build.list.ts";
import {OnInit} from "angular2/core";

@Component({
    template: require('./build.log.html')
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
