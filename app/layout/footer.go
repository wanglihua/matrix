package layout

import (
	"bytes"
	"github.com/revel/revel"
)

func GetFooter() string {
	var footer_buffer bytes.Buffer
	if revel.DevMode {
		footer_buffer.WriteString(`<script type="text/javascript" src="/static/js/vue.js"></script>`)
	} else {
		footer_buffer.WriteString(`<script type="text/javascript" src="/static/js/vue.min.js"></script>`)
	}
	footer_buffer.WriteString(`
        <script type="text/javascript" src="/static/ace/js/all1.js"></script>
        <!--[if lte IE 8]>
        <script type="text/javascript" src="/static/ace/js/html5shiv.min.js"></script>
        <script type="text/javascript" src="/static/ace/js/respond.min.js"></script>
        <![endif]-->
        <script type="text/javascript" src="/static/ace/js/all2.js"></script>
        <script type="text/javascript" src="/static/js/js.cookie.js"></script>
        <script type="text/javascript" src="/static/site.js"></script>
        <script type="text/javascript">
            try {
                ace.settings.check('breadcrumbs', 'fixed')
            } catch (e) {
            }
        </script>
    </div>
</div>
</body>
</html>
`)
	return footer_buffer.String()
}
