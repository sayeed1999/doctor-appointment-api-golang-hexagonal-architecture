import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DoctorsListComponent } from './doctors.component';

describe('DoctorsListComponent', () => {
  let component: DoctorsListComponent;
  let fixture: ComponentFixture<DoctorsListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DoctorsListComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DoctorsListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
