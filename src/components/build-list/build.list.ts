import {
    Component,
    Inject
} from "@angular/core";

import {
    Http,
    Headers
} from "@angular/http";

import {
    RouteParams,
    ROUTER_DIRECTIVES
} from "@angular/router-deprecated";

import * as moment from "moment";

import {Navigation} from "./../mdl-nav/mdl.nav";
import {BuildLog} from "./../build-log/build.log";

export class LogEntry {
    public id: string;
    public startTime: number;
    private _startTime: string;
    public duration: number;
    public state: string;
    public event: LogEvent;

    constructor(data) {
        this.id = data.id;
        this.startTime = data.start_time;
        this.duration  = data.duration;
        this.event     = new LogEvent(data.event);
        this.state = data.state;
    }

    static create(data) {
        return new LogEntry(data);
    }

    get fromnow(): string{
        return moment.unix(this.startTime).fromNow();
    }
}

export class LogEvent {
    public id: string;
    public type: string;
    public sender: LogEventSender;
    public description: string;

    constructor(data) {
        this.id = data.id;
        this.type = data.type;
        this.sender = new LogEventSender(data.sender);
        this.description = data.description;
    }
}

export class LogEventSender{
    public name: string;
    public avatar: string;
    public profile: string;

    constructor(data) {
        this.name = data.name;
        this.avatar = data.avatar_url;
        this.profile = data.profile_url;
    }
}

@Component({
    template: <string>require('./build.list.html'),
    directives: [ROUTER_DIRECTIVES, Navigation],
})
export class BuildList
{
    public builds: Array<LogEntry>;
    public repo: string;

    constructor(@Inject(Http) private http: Http, @Inject(RouteParams) public params: RouteParams) {
        this.builds = [];
        var hs = new Headers();
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        http.get("/api/builds/"+this.params.get("host")+"/"+this.params.get("namespace")+"/"+this.params.get("name"), {headers:hs})
            .map(res => res.json())
            .subscribe(val => {
                if (val) {
                    val.forEach((b) => {
                        this.builds.push(LogEntry.create(b));
                    })
                }
            });
    }
}
