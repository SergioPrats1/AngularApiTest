import { Component, OnInit, Inject, EventEmitter, Output } from '@angular/core';
import { FormGroup, FormControl, Validators, FormBuilder, ReactiveFormsModule } from '@angular/forms';
import { ApiReaderService } from '../api-reader/api-reader.service';
import { Router} from '@angular/router';

@Component({
  selector: 'app-add-song',
  templateUrl: './add-song.component.html',
  styleUrls: ['./add-song.component.css'],
  providers: [FormBuilder]
})
export class AddSongComponent implements OnInit {
  form: FormGroup;
  @Output() onClose: EventEmitter<boolean>;

  constructor( private formBuilder: FormBuilder, 
				private apiReaderService: ApiReaderService, private readonly router: Router) {
		this.onClose = new EventEmitter();					
	}

  ngOnInit(): void {
    this.form = this.formBuilder.group({ 
		  title: this.formBuilder.control('', Validators.compose([
		    Validators.required, Validators.pattern('[\\w\\-\\s\\/]+')]) ),
	    artist: this.formBuilder.control('', Validators.required),			
		  year: this.formBuilder.control('', this.yearValidator),
      comments: this.formBuilder.control('')
    });	  
  }
  
  yearValidator(control: FormControl) {
    if (control.value.trim().length === 0) {
      return null;
    }
    const year = parseInt(control.value, 10);
    const minYear = 1900;
    const maxYear = 2100;
    if (year >= minYear && year <= maxYear) {
      return null;
    } else {
      return {
        year: {
          min: minYear,
          max: maxYear
        }
      };
    }
  }
  
  
  onSubmit(newSong) {
    this.apiReaderService.addSong(newSong)
	      .subscribe(() => {    console.log("event to close the add form emitted ");
        this.router.navigate(['song-list']) /*this.onClose.emit(true)*/;
      });	  
  }

}
