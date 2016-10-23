package template

var MenuTemplate = `
{{$first := FirstEntity .entityList}}
            <li id="{{$first.ModuleLowerCase}}-menu" class="">
                <a href="#" class="dropdown-toggle">
                    <i class="menu-icon fa fa-users"></i>
                    <span class="menu-text">{{$first.ModuleChinese}}</span>
                    <b class="arrow fa fa-angle-down"></b>
                </a>
                <b class="arrow"></b>
                <ul class="submenu">
{{range $entityIndex, $entity := .entityList}}
                    <li id="{{$entity.ModuleLowerCase}}-{{LowerCase $entity.EntityCamelCase}}-menu" class="">
                        <a href="/{{$entity.ModuleLowerCase}}/{{LowerCase $entity.EntityCamelCase}}">
                            <i class="menu-icon fa fa-caret-right"></i>
                            {{$entity.EntityChinese}}
                        </a>
                        <b class="arrow"></b>
                    </li>
{{end}}
                </ul>
            </li>
`
