import { NgModule } from "@angular/core";
import { SharedModule } from "src/app/shared/shared.module";
import { AccountLoginComponent } from "./account-login/account-login.component";
import { AccountRegisterComponent } from "./account-register/account-register.component";
import { RoutingModule } from "./routing.module";

@NgModule({
    declarations: [
        AccountLoginComponent,
        AccountRegisterComponent,
    ],
    imports: [
        RoutingModule,
        SharedModule,
    ],
})
export class AuthModule{}