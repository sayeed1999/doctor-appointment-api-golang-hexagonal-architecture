import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: 'auth',
    loadChildren: () => import('./feature/auth/auth.module').then(x => x.AuthModule),
  },
  {
    path: '',
    loadChildren: () => import('./feature/home/home.module').then(x => x.HomeModule),
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
