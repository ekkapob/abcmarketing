$(function() {
  initServicesHover();
});

function initServicesHover() {
  var arrow = $(".service-details .arrow-up");
  var marketingDetails = $('.service-details ul.marketing');
  var developmentDetails = $('.service-details ul.development');

  var developmentCls = "development-arrow";
  var marketingCls = "marketing-arrow";
  $('#online-marketing').mouseover(function(){
    if (!arrow.hasClass(marketingCls)) {
      $(arrow).removeClass(developmentCls).addClass(marketingCls);
      $(marketingDetails).css('display', 'flex');
      $(developmentDetails).hide();
    }
  });
  $('#app-development').mouseover(function(){
    if (!arrow.hasClass(developmentCls)) {
      $(arrow).removeClass(marketingCls).addClass(developmentCls);
      $(marketingDetails).hide();
      $(developmentDetails).css('display', 'flex');
    }
  });
}
