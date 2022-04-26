import { Component, DoCheck, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.css']
})
export class ToolbarComponent implements OnInit {
  @Input() doctors: any[] = []

  // gorm.Model parameters are ID, CreatedAt, UpdatedAt, DeletedAt

  constructor(
    private router: Router
  ) { }

  ngOnInit(): void {

  }

  // ngDoCheck() {
  //   console.log(this.doctors)
  // }

  // navigateToDoctor(id: number) {
  //   this.router.navigate(['doctors', id])
  // }
}
