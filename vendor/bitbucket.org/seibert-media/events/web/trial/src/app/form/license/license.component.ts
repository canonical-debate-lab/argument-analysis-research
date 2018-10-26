import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';

@Component({
  selector: 'app-license',
  templateUrl: './license.component.html',
  styleUrls: ['./license.component.scss'],
})
export class LicenseComponent implements OnInit {

  form: FormGroup;

  constructor(
    fb: FormBuilder,
  ) {
    this.form = fb.group({
      'license': ['', [Validators.required]],
      'size': ['', [Validators.required]],
    });
  }

  ngOnInit() {
  }

}
