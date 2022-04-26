import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MakeAppointmentComponent } from './make-appointment.component';

describe('MakeAppointmentComponent', () => {
  let component: MakeAppointmentComponent;
  let fixture: ComponentFixture<MakeAppointmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MakeAppointmentComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MakeAppointmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
