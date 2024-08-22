export default function Infobox({text, number, color}) {
    const mystyle=`p-4 bg-${color}-100 rounded-xl text-gray-800 drop-shadow-lg`;
    return (
        
       <div class={mystyle}>
        <div class="font-bold text-2xl leading-none">{number}</div>
        <div class="mt-2">{text}</div>
      </div>
    );
};