import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MySongListComponent } from './my-song-list.component';

describe('MySongListComponent', () => {
  let component: MySongListComponent;
  let fixture: ComponentFixture<MySongListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MySongListComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MySongListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
