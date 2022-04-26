import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { BrowserAnimationsModule } from "@angular/platform-browser/animations";
import { RouterModule } from "@angular/router";
import { SharedModule } from "../shared/shared.module";
import { UtilityModule } from "../utility/utility.module";
import { AppLoaderComponent } from "./components/app-loader/app-loader.component";
import { AppShellComponent } from "./components/app-shell/app-shell.component";
import { ToolbarComponent } from './components/app-shell/toolbar/toolbar.component';

@NgModule({
    declarations: [
        AppShellComponent,
        AppLoaderComponent,
        ToolbarComponent,
    ],
    imports: [
        BrowserModule,
        BrowserAnimationsModule,
        UtilityModule,
        SharedModule,
        RouterModule,
    ],
    exports: [
        AppShellComponent,
        AppLoaderComponent,
    ],
    providers: [

    ]
})
export class CoreModule{}