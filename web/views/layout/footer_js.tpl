<!-- jQuery 3 -->
<script src="/static/js/admin/jquery.min.js"></script>
<!-- Bootstrap 3.3.7 -->
<script src="/static/js/admin/bootstrap.min.js"></script>
<!-- FastClick -->
<script src="/static/js/admin/fastclick.js"></script>


<!-- AdminLTE App -->
<script src="/static/js/admin/adminlte.min.js"></script>


<script>
$(document).ready(function(){
    /** add active class and stay opened when selected */
    var url = window.location;

    // for sidebar menu entirely but not cover treeview
    $('ul.sidebar-menu a').filter(function() {
         return this.href == url;
    }).parent().addClass('active');

    // for treeview
    $('ul.treeview-menu a').filter(function() {
         return this.href == url;
    }).parentsUntil(".sidebar-menu").addClass('active');
});
</script>

