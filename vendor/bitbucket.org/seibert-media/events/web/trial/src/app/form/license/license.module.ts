import {NgModule} from '@angular/core';
import {RouterModule} from '@angular/router';

import {LicenseComponent} from './license.component';
import {TranslateModule} from '@ngx-translate/core';
import {MaterialModule} from '../../material.module';

@NgModule({
  declarations: [LicenseComponent],
  imports: [
    MaterialModule,
    TranslateModule,
    RouterModule,
  ],
  exports: [LicenseComponent],
})
export class LicenseModule {
}
