import { CommonModule } from "@angular/common";
import { HttpClientModule } from "@angular/common/http";
import { NgModule } from "@angular/core";
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { BrowserModule } from "@angular/platform-browser";
import { BrowserAnimationsModule } from "@angular/platform-browser/animations";
import { AngularMaterialModule } from "./angular-material/angular-material.module";

@NgModule({
    declarations: [

    ],
    imports: [
        CommonModule,
        AngularMaterialModule,
        FormsModule,
        ReactiveFormsModule,
        HttpClientModule,
    ],
    exports: [
        CommonModule,
        AngularMaterialModule,
        FormsModule,
        ReactiveFormsModule,
        HttpClientModule,
    ],
    providers: [

    ]
})
export class SharedModule{}