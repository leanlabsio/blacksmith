import {
    Component,
    Input,
    OnInit,
    Inject
} from "@angular/core";

import {
    Headers,
    Http
} from "@angular/http";

import {
    Project,
    Repository,
    DockerExecutor,
    Image,
    Env
} from "./../dashboard/dashboard";

import {Observable} from "rxjs/Observable";
import {FORM_DIRECTIVES} from "@angular/common";
import {RouteParams} from "@angular/router-deprecated";
import {MdInput} from "./../mdl-textfield/mdl.textfield";

@Component({
    selector: "job-form",
    template: <string>require('./job.form.html'),
    directives: [FORM_DIRECTIVES, MdInput],
})
export class JobForm implements OnInit
{
    job: Project;

    constructor(@Inject(Http) private http: Http, @Inject(RouteParams) private params: RouteParams) {
        let hs = new Headers();
        hs.append("Authorization", "Bearer "+localStorage.getItem("jwt"));
        this.http.get("/api/projects/"+this.params.get("host")+"/"+this.params.get("namespace")+"/"+this.params.get("name"), {headers:hs})
            .map(res => <Project>res.json())
            .subscribe(job => this.job = job);
    }

    ngOnInit() {
        let image: Image = {}
        let repo: Repository = {clone_url: ""};
        let env: Env[] = [];
        let builder: DockerExecutor = {image: image, env: env};
        this.job = <Project>({executor: builder, repository: repo});
    }

    addenv() {
        console.log('asasdas')
        if (!this.job.executor.env || !this.job.executor.env.length) {
            let env: Env[] = [];
            this.job.executor.env = env;
        }
        console.log(this.job)
        this.job.executor.env.push(<Env>{name:"", value:""});
    }

    save() {
        var hs = new Headers();
        hs.append("Authorization", "Bearer "+localStorage.getItem("jwt"));
        console.log(this.job);
        this.http.put("/api/projects/"+this.params.get("host")+"/"+this.params.get("namespace")+"/"+this.params.get("name"), JSON.stringify(this.job), {headers:hs})
            .map(res => res.json())
            .subscribe(val => console.log(val));
    }
}
