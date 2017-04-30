import {
    Component,
    Inject,
    OnInit
} from "@angular/core";

import {
    Http,
    Headers,
    Response
} from "@angular/http";

import {RouteParams} from "@angular/router-deprecated";

import {LogEntry} from "./../build-list/build.list.ts";
import {NAVIGATION_DIRECTIVES} from "./../mdl-nav/mdl.nav";

const template: string = <string>require('./build.log.html');

@Component({
    template: template,
    directives: [NAVIGATION_DIRECTIVES]
})
export class BuildLog implements OnInit
{
    public build:string;

    constructor(@Inject(Http) private http: Http, @Inject(RouteParams) private params: RouteParams) {
        let host = params.get("host");
        let ns = params.get("namespace");
        let name = params.get("name");
        let commit = params.get("commit");
        let timestamp = params.get("timestamp");

        let hs = new Headers();
        hs.append("Authorization", "Bearer "+localStorage.getItem("jwt"));
        this.http.get("/api/logs/"+host+"/"+ns+"/"+name+"/"+commit+"/"+timestamp, {headers:hs})
            .map(res => res.json())
            .subscribe(val => this.build = val);
    }

    ngOnInit() {
        this.build = "";
    }
}
